package service

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"social_network_for_programmers/internal/entity"
	"social_network_for_programmers/internal/repository"
	"social_network_for_programmers/pkg/auth"
	"social_network_for_programmers/pkg/validationUserAuth"
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
	user := new(entity.UsersSignUpInput)
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		c.Writer.Write([]byte("Failed to read a user data. Please try again later."))
		log.Println("failed to deserialize json: ", err.Error())
		return
	}

	if err := validationUserAuth.ValidationUserSignUp(user); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.Writer.Write([]byte(err.Error() + ". Please try again."))
		log.Println("user is invalid: ", err.Error())
		return
	}

	hash := md5.Sum([]byte(user.Password))
	user.Password = hex.EncodeToString(hash[:])

	if err := h.repo.CreateUser(user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		c.Writer.Write([]byte("Failed to create user. Please try again later."))
		log.Println("failed to create a user: ", err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *AuthenticationService) SignInPage(c *gin.Context) {

}

func (h *AuthenticationService) SignIn(c *gin.Context) {
	user := entity.UsersSignInInput{}
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		c.Writer.Write([]byte("Failed to read a user data. Please try again later."))
		log.Println(err.Error())
		return
	}

	if user.Login == "" || user.Password == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.Writer.Write([]byte("Data is invalid. Please fill in all fields."))
		return
	}

	hash := md5.Sum([]byte(user.Password))
	hashPassword := hex.EncodeToString(hash[:])

	if err := h.repo.CheckUser(user.Login, hashPassword); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.Writer.Write([]byte("The user with this login or password was not found."))
		return
	}

	token, err := h.tokenManager.NewJwtToken(user.Login)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		c.Writer.Write([]byte("Please try again later."))
		log.Println("failed to create token: ", err.Error())
		return
	}

	c.Status(http.StatusOK)
	c.Writer.Write([]byte(token))
}
