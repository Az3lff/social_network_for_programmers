package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"social_network_for_programmers/internal/entity"
	"social_network_for_programmers/internal/repository"
)

type AuthenticationService struct {
	repo repository.Authentication
}

func NewAuthenticationService(repo repository.Authentication) *AuthenticationService {
	return &AuthenticationService{repo}
}

func (h *AuthenticationService) CreateUser(c *gin.Context) {
	user := entity.User{}
	if err := c.BindJSON(&user); err != nil {
		fmt.Println("Не удалось получить пользователя")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Проверяем, что пользователь с таким email еще не зарегистрирован
	//if _, exists := h.Storage.Users[user.ID]; exists {
	//	fmt.Println(errors.New("the user already exists"))
	//	c.AbortWithStatus(401)
	//	return
	//}
	//
	//h.Storage.Users[user.ID] = user

	c.Status(http.StatusOK)
}
