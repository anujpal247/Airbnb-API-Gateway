package main

import (
	"AuthApp/app"
	config "AuthApp/config/env"
)

func main() {
	config.Load() // load env variables from .env file

	cfg := app.NewConfig() // set the server
	app := app.NewApplication(cfg)

	app.Run()
}
