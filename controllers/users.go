package controllers

import (
	"AuthApp/dto"
	"AuthApp/services"
	"AuthApp/util"
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

func (uc *UserController) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {

	payload := r.Context().Value("payload").(*dto.SignupUserRequestDTO)
	fmt.Println("payload recived: ", payload)

	err := uc.UserService.CreateUser(payload)
	if err != nil {
		util.WriteJsonErrorResponse(w, http.StatusBadRequest, "something went worng", err)
		return
	}
	util.WriteJsonSuccessResponse(w, http.StatusOK, "User Registered successfully", true)
}

func (uc *UserController) LoginUserHandler(w http.ResponseWriter, r *http.Request) {

	payload := r.Context().Value("payload").(dto.LoginUserRequestDTO)

	jwtToken, err := uc.UserService.LoginUser(&payload)

	if jwtToken == "worng password" {
		util.WriteJsonErrorResponse(w, http.StatusBadRequest, "Worng Password", err)
		return
	}

	if err != nil {
		util.WriteJsonErrorResponse(w, http.StatusInternalServerError, "", err)
		return
	}

	util.WriteJsonSuccessResponse(w, http.StatusOK, "user logged is successfully", jwtToken)
}

func (uc *UserController) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(string)

	fmt.Println("user id from context", userId)

	if userId == "" {
		util.WriteJsonErrorResponse(w, http.StatusBadRequest, "user Id is required", fmt.Errorf("user id missing"))
		return
	}
	user, err := uc.UserService.GetUserById(userId)

	if err != nil {
		util.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to fetch user", err)
	}

	if user == nil {
		util.WriteJsonErrorResponse(w, http.StatusNotFound, "User not found", fmt.Errorf("user with id %s not found", userId))
		return
	}

	util.WriteJsonSuccessResponse(w, http.StatusOK, "user fetched successfully", user)
}
