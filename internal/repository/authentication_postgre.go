package repository

import (
	"errors"
	"fmt"
	"social_network_for_programmers/internal/entity"
)

var users []entity.User

type AuthenticationRepo struct {
	users map[int]entity.User
}

func NewAuthenticationRepo() *AuthenticationRepo {
	return &AuthenticationRepo{}
}

func (a *AuthenticationRepo) CreateUser(user *entity.User) error {
	users = append(users, *user)
	fmt.Println(users)
	return nil
}

func (a *AuthenticationRepo) CheckUser(login, password string) error {
	for _, userFromDB := range users {
		if userFromDB.Login == login && userFromDB.Password == password {
			return nil
		}
	}
	return errors.New("incorrect login or password")
}
