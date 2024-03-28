package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"social_network_for_programmers/internal/entity/responses"
	"social_network_for_programmers/internal/entity/users"
	"social_network_for_programmers/pkg/auth/utils"
)

func (h *Handler) initAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", h.userSignUp)
		auth.POST("/sign-in", h.userSignIn)
		auth.POST("/restore-access", h.userRestoreAccount)
		auth.POST("/check-restore-code", h.checkRestoreCode)
	}
}

func (h *Handler) userSignUp(c *gin.Context) {
	user := new(users.UserSignUp)
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "failed to read a user data, please try again"})
		log.Println("failed to deserialize json: ", err.Error())
		return
	}

	if err := utils.ValidationUserSignUp(user); err != nil {
		errResp := fmt.Sprintf("%s, please try again", err.Error())
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: errResp})
		return
	}

	if err := h.services.Auth.SignUp(c.Request.Context(), user); err != nil {
		errResp := fmt.Sprintf("failed to create user: %s", err.Error())
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: errResp})
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

	token, err := h.services.Auth.SignIn(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: err.Error()})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, responses.LoginResponse{AccessToken: token})
}

func (h *Handler) userRestoreAccount(c *gin.Context) {
	user := new(users.UserRestore)
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "failed to read a user data, please try again"})
		return
	}

	if !utils.EmailIsValid(user.Email) {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "incorrect email"})
		return
	}

	if err := h.services.Auth.RestoreAccount(c.Request.Context(), user.Email, &h.cfg.AuthEmail); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) checkRestoreCode(c *gin.Context) {
	req := new(users.RestoreAccessRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "failed to read a user data, please try again"})
		return
	}
	fmt.Println(req.Code)

	if !utils.CodeIsValid(req.Code) {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "incorrect code"})
		return
	}

	if err := h.services.Auth.CheckRestoreCode(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.InfoResponse{Message: "Access granted"})
}
