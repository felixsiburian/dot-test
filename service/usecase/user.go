package usecase

import (
	"dot-test/service"
	"dot-test/service/model"
	"dot-test/service/tools"
	"errors"
)

type userUsecase struct {
	userRepo service.IUserRepository
}

func (u userUsecase) UpdateEmail(email string, id string) error {
	if id == "" || email == "" {
		err := errors.New("invalid request")
		return tools.Wrap(err)
	}

	return u.userRepo.UpdateEmail(email, id)
}

func (u userUsecase) RetrieveById(id string) (*model.User, error) {
	if id == "" {
		err := errors.New("invalid request")
		return nil, tools.Wrap(err)
	}

	return u.userRepo.FindById(id)
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
) service.IUserUsecase {
	return userUsecase{
		userRepo: userRepo,
	}
}
