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
	q := `INSERT INTO users (login, email, hash_password) VALUES($1, $2, $3)`

	if _, err := u.db.Exec(ctx, q, user.Login, user.Email, user.Password); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			return errors.New(pgErr.Detail)
		}
		return err
	}

	return nil
}

func (u *UsersRepo) Find(ctx context.Context, user *users.UserSignIn) (userId string, err error) {
	q := `SELECT user_id FROM users WHERE email=$1 AND hash_password=$2`

	if err = u.db.QueryRow(ctx, q, user.Email, user.Password).Scan(&userId); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			err = fmt.Errorf("sql error: %s, details: %s, where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			return
		}
		return
	}

	return
}
