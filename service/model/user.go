package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID `gorm:"primary_key" json:"id"`
	Email       string    `json:"email" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Phonenumber string    `json:"phonenumber" validate:"required"`
	Username    string    `json:"username" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	Createdat   time.Time `json:"createdat"`
	Updatedat   time.Time `json:"updatedat"`
	Deletedat   time.Time `json:"deletedat"`
}
