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

	dbConfig := app.GetDBConfig()
	dbConn := db.ConnectionGorm(dbConfig)

	fmt.Println("dbConn: ", dbConn)
}
