package main

import (
	"AuthApp/app"
	dbConfig "AuthApp/config/db"
	config "AuthApp/config/env"
)

func main() {
	config.Load() // load env variables from .env file

	cfg := app.NewConfig() // set the server
	app := app.NewApplication(cfg)

	dbConfig.SetupDB()
	app.Run()
}
