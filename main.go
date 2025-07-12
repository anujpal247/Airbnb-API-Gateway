package main

import (
	"AuthApp/app"
)

func main() {
	cfg := app.NewConfig(":4000") // set the server to listen on port 4000
	app := app.NewApplication(cfg)

	app.Run()
}
