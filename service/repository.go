package service

import (
	"dot-test/service/model"
)

type IUserRepository interface {
	Create(payload model.User) (err error)
	UpdatePassword(password, id string) (err error)
	FindById(id string) (*model.User, error)
	Update(payload model.User) (err error)
}
