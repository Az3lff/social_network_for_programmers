package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"social_network_for_programmers/internal/entity"
)

type AuthenticationRepo struct {
	db *pgxpool.Pool
}

func NewAuthenticationRepo(db *pgxpool.Pool) *AuthenticationRepo {
	return &AuthenticationRepo{db}
}

func (a *AuthenticationRepo) CreateUser(user *entity.UsersSignUpInput) error {
	_, err := a.db.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, login VARCHAR(255) NOT NULL, email VARCHAR(255) NOT NULL, password VARCHAR(255) NOT NULL)")
	if err != nil {
		fmt.Println(err.Error())
	}

	row := a.db.QueryRow(context.Background(), "INSERT INTO users (login, email, password) VALUES ($1, $2, $3) RETURNING id", user.Login, user.Email, user.Password)
	var id uint64
	err = row.Scan(&id)

	if err != nil {
		log.Printf("failed to create user: %s", err.Error())
	}

	return nil
}

func (a *AuthenticationRepo) CheckUser(login, password string) error {

	return errors.New("incorrect login or password")
}
