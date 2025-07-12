package app

import (
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string // Port
}

type Application struct {
	Config Config
}

// constructor for config
func NewConfig(addr string) Config {
	return Config{
		Addr: addr,
	}
}

// constructor for application
func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}

func (app *Application) Run() error {
	server := http.Server{
		Addr:         app.Config.Addr,
		Handler:      nil, // setup chi router
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server is running on", app.Config.Addr)

	return server.ListenAndServe()
}
