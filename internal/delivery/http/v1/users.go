package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"social_network_for_programmers/internal/entity/users"
	"social_network_for_programmers/pkg/auth"
	"social_network_for_programmers/pkg/responses"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		user.POST("/sign-up", h.userSignUp)
		user.POST("/sign-in", h.userSignIn)

		user.GET("/sign-up")
		user.GET("/sign-in")
	}
}

func (h *Handler) userSignUp(c *gin.Context) {
	user := new(users.UserSignUp)
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "failed to read a user data, please try again"})
		log.Println("failed to deserialize json: ", err.Error())
		return
	}

	if err := auth.ValidationUserSignUp(user); err != nil {
		errResp := fmt.Sprintf("%s, please try again", err.Error())
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: errResp})
		return
	}

	if err := h.services.Users.SignUp(c.Request.Context(), user); err != nil {
		errResp := fmt.Sprintf("failed to create user: %s", err.Error())
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: errResp})
		log.Println(errResp)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) userSignIn(c *gin.Context) {
	user := new(users.UserSignIn)
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "failed to read a user data, please try again"})
		log.Println(err.Error())
		return
	}

	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "data is invalid, please fill in all fields"})
		return
	}

	token, err := h.services.Users.SignIn(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: err.Error()})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, responses.LoginResponse{AccessToken: token})
}
