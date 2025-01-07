package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type ConfigDB struct {
	Driver   string
	DBName   string
	Username string
	Password string
	Host     string
	Port     string
}

func ConnectionGorm(c ConfigDB) *gorm.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		c.Host, c.Port, c.Username, c.DBName, "disable", c.Password)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
