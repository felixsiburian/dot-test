package service

import (
	"dot-test/service/model"
)

type IUserRepository interface {
	Create(payload model.User) (err error)
	UpdateEmail(email, id string) (err error)
	FindById(id string) (*model.User, error)
}
