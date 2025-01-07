package service

import (
	"dot-test/service/model"
)

type IUserRepository interface {
	Create(payload model.User) (err error)
}
