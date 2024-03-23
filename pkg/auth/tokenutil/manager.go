package tokenutil

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenManager interface {
	NewJwtToken(login string) (string, error)
}

type Manager struct {
	key string
}

func NewManager(key string) (*Manager, error) {
	if key == "" {
		return nil, errors.New("empty key")
	}

	return &Manager{key}, nil
}

func (m *Manager) NewJwtToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(m.key))
}
