package utils

import "golang.org/x/crypto/bcrypt"

func CompareHashPassword(inputPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}
