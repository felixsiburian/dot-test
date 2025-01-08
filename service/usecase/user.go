package usecase

import (
	"dot-test/service"
	"dot-test/service/model"
	"dot-test/service/tools"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
)

type userUsecase struct {
	userRepo    service.IUserRepository
	redisClient *redis.Client
}

var (
	userKey = "user-"
)

func (u userUsecase) UpdateEmail(email string, id string) error {
	if id == "" || email == "" {
		err := errors.New("invalid request")
		return tools.Wrap(err)
	}

	redisKey := userKey + id
	userData, err := u.redisClient.Get(redisKey).Result()
	if err != nil && err != redis.Nil {
		return tools.Wrap(err)
	}

	if userData != "" {
		u.redisClient.Del(redisKey)
	}

	if err := u.userRepo.UpdateEmail(email, id); err != nil {
		return tools.Wrap(err)
	}

	user, err := u.userRepo.FindById(id)
	if err != nil {
		return tools.Wrap(err)
	}

	userStr, _ := json.Marshal(user)
	u.redisClient.Set(redisKey, string(userStr), 0)

	return nil
}

func (u userUsecase) RetrieveById(id string) (*model.User, error) {
	var response *model.User
	if id == "" {
		err := errors.New("invalid request")
		return nil, tools.Wrap(err)
	}

	redisKey := userKey + id
	userData, err := u.redisClient.Get(redisKey).Result()
	if err != nil && err != redis.Nil {
		return nil, tools.Wrap(err)
	}

	if userData == "" {
		res, err := u.userRepo.FindById(id)
		if err != nil {
			return nil, tools.Wrap(err)
		}

		resStr, _ := json.Marshal(res)

		resp := u.redisClient.Set(redisKey, string(resStr), 0)
		if resp.Err() != nil {
			return nil, tools.Wrap(resp.Err())
		}

		return res, nil
	}

	if err := json.Unmarshal([]byte(userData), &response); err != nil {
		return nil, tools.Wrap(err)
	}

	return response, nil
}

func (u userUsecase) Create(request model.User) (err error) {
	password, err := tools.HashPassword(request.Password)
	if err != nil {
		return tools.Wrap(err)
	}

	if err = tools.Validate(request); err != nil {
		return tools.Wrap(err)
	}

	request.Password = password
	return u.userRepo.Create(request)
}

func NewUserUsecase(
	userRepo service.IUserRepository,
	redisClient *redis.Client,
) service.IUserUsecase {
	return userUsecase{
		userRepo:    userRepo,
		redisClient: redisClient,
	}
}
