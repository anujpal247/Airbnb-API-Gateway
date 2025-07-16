package services

import (
	db "AuthApp/db/repositories"
	"fmt"
)

type UserService interface {
	CreateUser() error
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
