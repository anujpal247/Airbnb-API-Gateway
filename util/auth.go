package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error hassing password ", err)
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(plainPassword string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

	if err != nil {
		fmt.Println("worng password")
		return err
	}
	return nil
}
