package app

import (
	dbConfig "AuthApp/config/db"
	config "AuthApp/config/env"
	"AuthApp/controllers"
	repo "AuthApp/db/repositories"
	"AuthApp/router"
	"AuthApp/services"
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
func NewConfig() Config {
	port := config.GetString("PORT", ":8080")

	return Config{
		Addr: port,
	}
}

// constructor for application
func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}

func (app *Application) Run() error {

	db, err := dbConfig.SetupDB()

	if err != nil {
		fmt.Println("Error setting up db", err)
		return err
	}

	ur := repo.NewUserRepository(db)
	us := services.NewUserService(ur)
	uc := controllers.NewUserController(us)
	uRouter := router.NewUserRouter(uc)

	server := http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(uRouter), // setup chi router
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server is running on", app.Config.Addr)

	return server.ListenAndServe()
}
