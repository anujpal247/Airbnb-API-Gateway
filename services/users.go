package services

import (
	env "AuthApp/config/env"
	db "AuthApp/db/repositories"
	"AuthApp/dto"
	"AuthApp/util"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type UserService interface {
	CreateUser(payload *dto.SignupUserRequestDTO) error
	LoginUser(payload *dto.LoginUserRequestDTO) (string, error)
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) CreateUser(payload *dto.SignupUserRequestDTO) error {
	password := payload.Password
	email := payload.Email
	username := payload.Username

	hashPW, err := util.HashPassword(password)
	if err != nil {
		fmt.Println("Error hassing pw", err)
		return err
	}

	u.userRepository.Create(username, email, hashPW)
	return nil
}

func (u *UserServiceImpl) LoginUser(payload *dto.LoginUserRequestDTO) (string, error) {

	email := payload.Email
	password := payload.Password

	fmt.Println("Getting user in user service")
	user, err := u.userRepository.GetByEmail(email)
	if err != nil {
		fmt.Println("Error fetching user by email", err)
		return "", err
	}
	fmt.Println("Response ", user)
	isPasswordValid := util.CheckPasswordHash(password, user.Password)

	if !isPasswordValid {
		fmt.Println("worng password")
		return "", nil
	}

	jwtPayload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)

	tokenString, tokenErr := token.SignedString([]byte(env.GetString("JWT_TOKEN", "TOKEN")))

	if tokenErr != nil {
		fmt.Println("Error signin token", tokenErr)
		return "", tokenErr
	}
	return tokenString, nil
}
