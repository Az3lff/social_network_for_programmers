package authentication

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"social_network_for_programmers/internal/service/authentication/models"
)

type AuthHandler struct {
	Storage *AuthStorage
}

func GetAuthenticationPage(c *gin.Context) {

}

func (h *AuthHandler) Register(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("Не удалось получить пользователя")
		c.AbortWithStatus(401)
		return
	}

	// Проверяем, что пользователь с таким email еще не зарегистрирован
	if _, exists := h.Storage.Users[user.ID]; exists {
		fmt.Println(errors.New("the user already exists"))
		c.AbortWithStatus(401)
		return
	}

	h.Storage.Users[user.ID] = user

	c.AbortWithStatus(200)
	fmt.Println(h.Storage.Users)
}
