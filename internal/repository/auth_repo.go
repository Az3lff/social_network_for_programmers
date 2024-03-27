package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"social_network_for_programmers/internal/entity/users"
)

type AuthRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{db}
}

func (a *AuthRepo) Create(ctx context.Context, user *users.UserSignUp) error {
	q := `INSERT INTO users (login, email, hash_password) VALUES($1, $2, $3)`

	if _, err := a.db.Exec(ctx, q, user.Login, user.Email, user.Password); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			return errors.New(pgErr.Detail)
		}
		return err
	}

	return nil
}

func (a *AuthRepo) GetByEmail(ctx context.Context, email string) (*users.UserRepo, error) {
	q := `SELECT user_id, login, email, hash_password FROM users WHERE email=$1`

	userRepo := new(users.UserRepo)
	if err := a.db.QueryRow(ctx, q, email).Scan(&userRepo.UUID, &userRepo.Login, &userRepo.Email, &userRepo.HashPassword); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			err = fmt.Errorf("sql error: %s, details: %s, where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			return nil, err
		}
		return nil, err
	}

	return userRepo, nil
}

func (a *AuthRepo) FindByEmail(ctx context.Context, email string) error {
	q := `SELECT COUNT(*) FROM users WHERE email=$1`

	var count int
	if err := a.db.QueryRow(ctx, q, email).Scan(&count); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			err = fmt.Errorf("sql error: %s, details: %s, where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			return err
		}
		return err
	}

	if count < 1 {
		return errors.New("user with such an emails does not exist")
	}

	return nil
}
