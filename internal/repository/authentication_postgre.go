package repository

import "github.com/gin-gonic/gin"

type AuthenticationRepo struct {
}

func NewAuthenticationRepo() *AuthenticationRepo {
	return &AuthenticationRepo{}
}

func (a *AuthenticationRepo) CreateUser(c *gin.Context) {

}
