package service

import (
	"dot-test/service/model"
)

type IUserUsecase interface {
	Create(request model.User) (err error)
	RetrieveById(id string) (*model.User, error)
	UpdatePassword(password string, id string) error
	Update(user model.User) error
}
