package repository

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"social_network_for_programmers/internal/entity"
)

var users []entity.UsersSignUpInput

type AuthenticationRepo struct {
	db *pgx.Conn
}

func NewAuthenticationRepo(db *pgx.Conn) *AuthenticationRepo {
	return &AuthenticationRepo{db}
}

func (a *AuthenticationRepo) CreateUser(user *entity.UsersSignUpInput) error {
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
