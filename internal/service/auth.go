package service

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"social_network_for_programmers/internal/config"
	"social_network_for_programmers/internal/entity/emails"
	"social_network_for_programmers/internal/entity/users"
	"social_network_for_programmers/internal/repository"
	"social_network_for_programmers/pkg/auth"
	"social_network_for_programmers/pkg/auth/utils"
)

type AuthService struct {
	repo         repository.Auth
	tokenManager auth.TokenManager
}

func NewAuthService(repo repository.Auth, tokenManager auth.TokenManager) *AuthService {
	return &AuthService{repo, tokenManager}
}

func (a *AuthService) SignUp(ctx context.Context, user *users.UserSignUp) error {
	hashPassword, err := utils.CreateHashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPassword

	if err := a.repo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (a *AuthService) SignIn(ctx context.Context, user *users.UserSignIn) (string, error) {
	userRepo, err := a.repo.GetByEmail(ctx, user.Email)
	if err != nil && err.Error() != pgx.ErrNoRows.Error() {
		return "", err
	}

	if err != nil && err.Error() == pgx.ErrNoRows.Error() {
		return "", errors.New("incorrect emails addresses")
	}

	if !utils.CompareHashPassword(user.Password, userRepo.HashPassword) {
		return "", errors.New("incorrect password")
	}

	token, err := a.tokenManager.NewJwtToken(userRepo.UUID)
	if err != nil {
		return "", errors.New("failed to create token: " + err.Error())
	}

	return token, nil
}

func (a *AuthService) RestoreAccount(ctx context.Context, email string, ath *config.AuthEmail) error {
	if err := a.repo.FindByEmail(ctx, email); err != nil {
		return err
	}

	code := utils.GenerateRestoreCode()

	content := "Subject:Restore account code\n" + code

	s := emails.MessageEmail{
		Ath:     ath,
		From:    ath.Username,
		To:      []string{email},
		Content: []byte(content),
	}

	if err := utils.SendMessageEmail(&s); err != nil {
		return err
	}

	return nil
}
