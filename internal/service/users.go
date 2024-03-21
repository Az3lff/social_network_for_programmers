package service

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"social_network_for_programmers/internal/entity/users"
	"social_network_for_programmers/internal/repository"
	"social_network_for_programmers/pkg/auth"
	"social_network_for_programmers/pkg/validationUserAuth"
)

type UsersService struct {
	repo         repository.Users
	tokenManager auth.TokenManager
}

func NewUsersService(repo repository.Users, tokenManager auth.TokenManager) *UsersService {
	return &UsersService{repo, tokenManager}
}

func (h *UsersService) SignUpPage(c *gin.Context) {

}

func (h *UsersService) SignUp(c *gin.Context) {
	user := new(users.UsersSignUpInput)
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

func (h *UsersService) SignInPage(c *gin.Context) {

}

func (h *UsersService) SignIn(c *gin.Context) {
	user := users.UsersSignInInput{}
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
