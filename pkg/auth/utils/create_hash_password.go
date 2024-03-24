package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func CreateHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hashing password: %s", err.Error())
	}

	return string(hash), nil
}
