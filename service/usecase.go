package service

import (
	"dot-test/service/model"
)

type IUserUsecase interface {
	Create(request model.User) (err error)
}
