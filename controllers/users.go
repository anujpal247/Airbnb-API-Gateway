package controllers

import (
	"AuthApp/dto"
	"AuthApp/services"
	"AuthApp/util"
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
	var payload dto.SignupUserRequestDTO

	jsonErr := util.ReadJsonBody(r, &payload)
	if jsonErr != nil {
		util.WriteJsonErrorResponse(w, http.StatusBadRequest, "Something went worng while reading json body", jsonErr)
		return
	}

	if validationErr := util.Validator.Struct(payload); validationErr != nil {
		util.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input data", validationErr)
		return
	}
	err := uc.UserService.CreateUser(&payload)
	if err != nil {
		util.WriteJsonErrorResponse(w, http.StatusBadRequest, "something went worng", err)
		return
	}
	util.WriteJsonSuccessResponse(w, http.StatusOK, "User Registered successfully", true)
}

func (uc *UserController) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload dto.LoginUserRequestDTO

	// jsonErr := util.ReadJsonBody(r, &payload)

	// if jsonErr != nil {
	// 	w.Write([]byte("something went worng"))
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	if jsonErr := util.ReadJsonBody(r, &payload); jsonErr != nil {
		util.WriteJsonErrorResponse(w, http.StatusBadRequest, "something went worng", jsonErr)
		return
	}

	if validationErr := util.Validator.Struct(payload); validationErr != nil {
		util.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input data", validationErr)
		return
	}

	jwtToken, err := uc.UserService.LoginUser(&payload)
	if err != nil {
		util.WriteJsonErrorResponse(w, http.StatusInternalServerError, "", err)
		return
	}

	util.WriteJsonSuccessResponse(w, http.StatusOK, "user logged is successfully", jwtToken)
}
