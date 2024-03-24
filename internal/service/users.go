package service

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"social_network_for_programmers/internal/entity/users"
	"social_network_for_programmers/internal/repository"
	"social_network_for_programmers/pkg/auth"
	"social_network_for_programmers/pkg/auth/utils"
)

type UsersService struct {
	repo         repository.Users
	tokenManager auth.TokenManager
}

func NewUsersService(repo repository.Users, tokenManager auth.TokenManager) *UsersService {
	return &UsersService{repo, tokenManager}
}

func (u *UsersService) SignUp(ctx context.Context, user *users.UserSignUp) error {
	hashPassword, err := utils.CreateHashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPassword

	if err := u.repo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *UsersService) SignIn(ctx context.Context, user *users.UserSignIn) (string, error) {
	userRepo, err := u.repo.GetByEmail(context.Background(), user.Email)
	if err != nil && err.Error() != pgx.ErrNoRows.Error() {
		return "", err
	}

	if err != nil && err.Error() == pgx.ErrNoRows.Error() {
		return "", errors.New("incorrect email addresses")
	}

	if !utils.CompareHashPassword(user.Password, userRepo.HashPassword) {
		return "", errors.New("incorrect password")
	}

	token, err := u.tokenManager.NewJwtToken(userRepo.UUID)
	if err != nil {
		return "", errors.New("failed to create token: " + err.Error())
	}

	return token, nil
}
