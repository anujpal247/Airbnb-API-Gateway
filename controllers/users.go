package controllers

import (
	"AuthApp/services"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register user is called in user controller")
	uc.UserService.CreateUser()
	w.Write([]byte("User registration endpoint"))
}

func (uc *UserController) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user fetching in controller layer")
	uc.UserService.GetUserById()
	w.Write([]byte("get user endpoint"))
}
