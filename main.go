package main

import (
	"dot-test/lib/db"
	"dot-test/service/config"
	"dot-test/service/delivery/router"
	"dot-test/service/repository"
	"dot-test/service/usecase"
	"fmt"
	"github.com/labstack/echo"
	"log"
	"os"
)

func main() {
	start()
}

func start() {
	app := config.Config{}
	e := echo.New()

	app.CatchError(app.InitEnv())

	//register db
	dbConfig := app.GetDBConfig()
	dbConn := db.ConnectionGorm(dbConfig)

	//register redis
	redisConfig := app.GetRedisConfig()
	redisConn := db.RedisConnection(redisConfig)
	fmt.Println("redisConn: ", redisConn)

	userRepo := repository.NewUserRepository(dbConn)

	userUsecase := usecase.NewUserUsecase(userRepo, redisConn)

	router.NewRouter(e, userUsecase)

	log.Println("service running on port: ", os.Getenv("APP_PORT"))
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))))
}
