package usecase

import (
	"dot-test/service"
	"dot-test/service/model"
	"dot-test/service/tools"
)

type userUsecase struct {
	userRepo service.IUserRepository
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
