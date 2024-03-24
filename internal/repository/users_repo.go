package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"social_network_for_programmers/internal/entity/users"
)

type UsersRepo struct {
	db *pgxpool.Pool
}

func NewUsersRepo(db *pgxpool.Pool) *UsersRepo {
	return &UsersRepo{db}
}

func (u *UsersRepo) Create(ctx context.Context, user *users.UserSignUp) error {
	q := `INSERT INTO users (login, email, utils) VALUES($1, $2, $3)`

	if _, err := u.db.Exec(ctx, q, user.Login, user.Email, user.Password); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			return errors.New(pgErr.Detail)
		}
		return err
	}

	return nil
}

func (u *UsersRepo) GetByEmail(ctx context.Context, email string) (*users.UserRepo, error) {
	q := `SELECT user_id, login, email, utils FROM users WHERE email=$1`

	userRepo := new(users.UserRepo)
	if err := u.db.QueryRow(ctx, q, email).Scan(&userRepo.UUID, &userRepo.Login, &userRepo.Email, &userRepo.HashPassword); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			err = fmt.Errorf("sql error: %s, details: %s, where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			return nil, err
		}
		return nil, err
	}

	return userRepo, nil
}
