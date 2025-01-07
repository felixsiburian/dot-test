package repository

import (
	"github.com/jinzhu/gorm"
	"log"
)

func ExecuteTransaction(db *gorm.DB, fn func(db2 *gorm.DB) error) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Fatalf("failed transaction: %v", r)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
