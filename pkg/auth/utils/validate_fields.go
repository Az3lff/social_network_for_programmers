package utils

import (
	"errors"
	"net/mail"
	"social_network_for_programmers/internal/entity/auth_entity"
	"strconv"
	"strings"
	"unicode"
)

func ValidationUserSignUp(user *auth_entity.UserSignUp) error {
	var errStr []string
	if !LoginIsValid(user.Login) {
		errStr = append(errStr, "login="+user.Login)
	}
	if !EmailIsValid(user.Email) {
		errStr = append(errStr, "emails="+user.Email)
	}
	if !PasswordIsValid(user.Password) {
		errStr = append(errStr, "password=*hidden*")
	}

	if len(errStr) > 0 {
		return errors.New(strings.Join(errStr, ", ") + " is incorrect")
	}

	return nil
}

func EmailIsValid(email string) bool {
	if len([]rune(email)) < 5 || len([]rune(email)) > 255 {
		return false
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return false
	}

	for _, char := range email {
		if char > unicode.MaxASCII {
			return false
		}
	}

	return true
}

func LoginIsValid(login string) bool {
	if len([]rune(login)) < 1 || len([]rune(login)) > 255 {
		return false
	}

	for _, char := range login {
		if char > unicode.MaxASCII {
			return false
		}
	}

	return true
}

func PasswordIsValid(password string) bool {
	if len([]rune(password)) < 1 || len([]rune(password)) > 255 {
		return false
	}

	for _, char := range password {
		if char > unicode.MaxASCII {
			return false
		}
	}

	return true
}

func CodeIsValid(code string) bool {
	if len([]rune(code)) != 6 {
		return false
	}

	_, err := strconv.Atoi(code)
	if err != nil {
		return false
	}

	return true
}
