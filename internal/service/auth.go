package service

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"social_network_for_programmers/internal/config"
	"social_network_for_programmers/internal/entity/auth_entity"
	"social_network_for_programmers/internal/entity/emails"
	"social_network_for_programmers/internal/repository"
	"social_network_for_programmers/pkg/auth"
	"social_network_for_programmers/pkg/auth/utils"
	"time"
)

type AuthService struct {
	repo         repository.Auth
	tokenManager auth.TokenManager
	cache        *redis.Client
}

func NewAuthService(repo repository.Auth, tokenManager auth.TokenManager, cache *redis.Client) *AuthService {
	return &AuthService{repo, tokenManager, cache}
}

func (a *AuthService) SignUp(ctx context.Context, user *auth_entity.UserSignUp) error {
	hashPassword, err := utils.GenerateHashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPassword

	if err := a.repo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (a *AuthService) SignIn(ctx context.Context, user *auth_entity.UserSignIn) (string, error) {
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
	a.cache.Set(ctx, email, code, 3*time.Minute)

	content := "Subject:Restore account code\n" + code

	s := &emails.MessageEmail{
		Ath:     ath,
		From:    ath.Username,
		To:      []string{email},
		Content: []byte(content),
	}

	if err := utils.SendMessageEmail(s); err != nil {
		return err
	}

	return nil
}

func (a *AuthService) CheckRestoreCode(ctx context.Context, req *auth_entity.RestoreAccessRequest) error {
	codeInCache, err := a.cache.Get(ctx, req.Email).Result()
	if err != nil {
		return err
	}

	if !utils.CompareRestoreCode(req.Code, codeInCache) {
		return errors.New("code is invalid")
	}

	a.cache.Del(ctx, req.Email)

	return nil
}

func (a *AuthService) UpdatePassword(ctx context.Context, user *auth_entity.UserUpdatePassword, ath *config.AuthEmail) error {
	hashPassword, err := utils.GenerateHashPassword(user.HashPassword)
	if err != nil {
		return err
	}
	user.HashPassword = hashPassword

	if err := a.repo.UpdatePasswordByEmail(ctx, user); err != nil {
		return err
	}

	content := "Subject:Restore account\n" + "The password has been changed"

	s := &emails.MessageEmail{
		Ath:     ath,
		From:    ath.Username,
		To:      []string{user.Email},
		Content: []byte(content),
	}

	if err := utils.SendMessageEmail(s); err != nil {
		return err
	}

	return nil
}
