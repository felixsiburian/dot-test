package repository

import (
	"dot-test/service"
	"dot-test/service/model"
	"dot-test/service/tools"
	"github.com/jinzhu/gorm"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func (u UserRepository) FindById(id string) (*model.User, error) {
	var res model.User
	err := u.db.Debug().Table("user").Select("id,email,name,phonenumber,username,createdat,updatedat,deletedat").Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil, tools.Wrap(err)
	}

	return &res, nil
}

func (u UserRepository) UpdateEmail(email, id string) (err error) {
	if err := u.db.Debug().Table("user").Where("id = ? ", id).Update("email", email).Error; err != nil {
		return tools.Wrap(err)
	}

	return nil
}

func (u UserRepository) Create(payload model.User) (err error) {
	payload.Createdat = time.Now()
	payload.Updatedat = time.Now()

	err = ExecuteTransaction(u.db, func(db2 *gorm.DB) error {
		if err := u.db.Debug().Table("user").Create(&payload).Error; err != nil {
			return tools.Wrap(err)
		}
		return nil
	})

	return err
}

func NewUserRepository(db *gorm.DB) service.IUserRepository {
	return UserRepository{
		db: db,
	}
}
