package main

import (
	"dot-test/lib/db"
	"dot-test/service/config"
	"fmt"
)

func main() {
	start()
}

func start() {
	app := config.Config{}
	//e := echo.New()

	app.CatchError(app.InitEnv())

	//register db
	dbConfig := app.GetDBConfig()
	dbConn := db.ConnectionGorm(dbConfig)

	//register redis
	redisConfig := app.GetRedisConfig()
	redisConn := db.RedisConnection(redisConfig)

	fmt.Println("dbConn: ", dbConn)
	fmt.Println("redisConn: ", redisConn)
}
