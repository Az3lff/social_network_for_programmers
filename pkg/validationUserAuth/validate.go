package validationUserAuth

import (
	"errors"
	"net/mail"
	"social_network_for_programmers/internal/entity"
	"strings"
	"unicode"
)

func ValidationUserSignUp(user *entity.UsersSignUpInput) error {
	var errStr []string
	if !loginIsValid(user.Login) {
		errStr = append(errStr, "login="+user.Login)
	}
	if !emailIsValid(user.Email) {
		errStr = append(errStr, "email="+user.Email)
	}
	if !passwordIsValid(user.Password) {
		errStr = append(errStr, "password=*hidden*")
	}

	if len(errStr) > 0 {
		return errors.New(strings.Join(errStr, ", ") + " is invalid")
	}

	return nil
}

func emailIsValid(email string) bool {
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

func loginIsValid(login string) bool {
	if len([]rune(login)) < 0 || len([]rune(login)) > 255 {
		return false
	}

	for _, char := range login {
		if char > unicode.MaxASCII {
			return false
		}
	}

	return true
}

func passwordIsValid(password string) bool {
	if len([]rune(password)) < 0 || len([]rune(password)) > 255 {
		return false
	}

	for _, char := range password {
		if char > unicode.MaxASCII {
			return false
		}
	}

	return true
}
