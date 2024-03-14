package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"social_network_for_programmers/internal/entity"
	"social_network_for_programmers/internal/repository"
	"social_network_for_programmers/pkg/auth"
)

type AuthenticationService struct {
	repo         repository.Authentication
	tokenManager auth.TokenManager
}

func NewAuthenticationService(repo repository.Authentication, tokenManager auth.TokenManager) *AuthenticationService {
	return &AuthenticationService{repo, tokenManager}
}

func (h *AuthenticationService) SignUpPage(c *gin.Context) {

}

func (h *AuthenticationService) SignUp(c *gin.Context) {
	user := entity.User{}
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(fmt.Errorf("failed to get user: %s", err.Error()))
		return
	}

	if user.Login == "" || user.Email == "" || user.Password == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		c.Writer.Write([]byte("all fields must be filled in"))
		return
	}
	hash := md5.Sum([]byte(user.Password))
	user.Password = hex.EncodeToString(hash[:])

	if err := h.repo.CreateUser(&user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Println(fmt.Errorf("failed to create user: %s", err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func (h *AuthenticationService) SignInPage(c *gin.Context) {

}

func (h *AuthenticationService) SignIn(c *gin.Context) {
	user := entity.User{}
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(fmt.Errorf("failed to get user: %s", err.Error()))
		return
	}

	if user.Login == "" || user.Password == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		c.Writer.Write([]byte("all fields must be filled in"))
		return
	}
	hash := md5.Sum([]byte(user.Password))
	hashPassword := hex.EncodeToString(hash[:])
	if err := h.repo.CheckUser(user.Login, hashPassword); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		c.Writer.Write([]byte(err.Error()))
		return
	}
	c.AbortWithStatus(http.StatusOK)
	c.Writer.Write([]byte("Hello " + user.Login + "!"))

	token, err := h.tokenManager.NewJwtToken(user.Login)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	fmt.Println(token)

	fmt.Println()
	fmt.Printf("ID : %d, Login: %s, Email: %s, Password: %s", user.ID, user.Login, user.Email, user.Password)
	fmt.Println()
}
