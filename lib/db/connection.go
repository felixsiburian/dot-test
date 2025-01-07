package db

import (
	"fmt"
	"github.com/go-redis/redis"
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

type RedisConfig struct {
	Host string
	Port string
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

func RedisConnection(c RedisConfig) *redis.Client {
	addr := fmt.Sprintf("%s:%s", c.Host, c.Port)
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return client
}
