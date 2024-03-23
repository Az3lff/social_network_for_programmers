package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/jackc/pgx/v5"
	"social_network_for_programmers/internal/entity/users"
	"social_network_for_programmers/internal/repository"
	"social_network_for_programmers/pkg/auth/tokenutil"
)

type UsersService struct {
	repo         repository.Users
	tokenManager tokenutil.TokenManager
}

func NewUsersService(repo repository.Users, tokenManager tokenutil.TokenManager) *UsersService {
	return &UsersService{repo, tokenManager}
}

func (u *UsersService) SignUp(ctx context.Context, user *users.UserSignUp) error {
	hash := md5.Sum([]byte(user.Password))
	user.Password = hex.EncodeToString(hash[:])

	if err := u.repo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *UsersService) SignIn(ctx context.Context, user *users.UserSignIn) (string, error) {
	hash := md5.Sum([]byte(user.Password))
	user.Password = hex.EncodeToString(hash[:])

	userId, err := u.repo.Find(context.Background(), user)
	if err != nil && err.Error() != pgx.ErrNoRows.Error() {
		return "", err
	}

	if err != nil && err.Error() == pgx.ErrNoRows.Error() {
		return "", errors.New("user with this login or password was not found")
	}

	token, err := u.tokenManager.NewJwtToken(userId)
	if err != nil {
		return "", errors.New("failed to create token: " + err.Error())
	}

	return token, nil
}
