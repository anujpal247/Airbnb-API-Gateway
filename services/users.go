package services

import (
	db "AuthApp/db/repositories"
	"AuthApp/models"
	"fmt"
)

type UserService interface {
	CreateUser() error
	GetUserById() (*models.User, error)
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("creating user in user service")
	u.userRepository.Create()
	return nil
}

func (u *UserServiceImpl) GetUserById() (*models.User, error) {
	fmt.Println("Getting user in user service")
	res, err := u.userRepository.GetById()
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}
	fmt.Println("Response ", res)
	return res, nil
}
