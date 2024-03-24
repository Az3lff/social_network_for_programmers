package tokenutil

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func IsAuthorized(requestToken, secretKey string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			newErr := fmt.Sprintf("unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New(newErr)
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
