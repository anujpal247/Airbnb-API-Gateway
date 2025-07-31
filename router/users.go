package router

import (
	"AuthApp/controllers"
	"AuthApp/middlewares"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController) Router {
	return &UserRouter{
		userController: _userController,
	}
}

func (ur *UserRouter) Register(r chi.Router) {
	r.With(middlewares.UserSignupRequestValidator).Post("/signup", ur.userController.RegisterUserHandler)
	r.With(middlewares.UserLoginRequestValidator).Get("/login", ur.userController.LoginUserHandler)
}
