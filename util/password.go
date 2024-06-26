package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password: %w", err)
	}
	return string(hashed_password), nil
}

func CheckPassword(password string, hashed_password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
	return err

}
