package v1

import (
	"social_network_for_programmers/internal/middleware/auth"
	"social_network_for_programmers/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services}
}

func (h *Handler) Init(api *gin.RouterGroup, secretKey string) {
	v1 := api.Group("/v1")
	{
		h.initUsersRoutes(v1)
		protectedRouter := v1.Group("")
		protectedRouter.Use(auth.JwtAuthMiddleware(secretKey))

		h.initMessengerRoutes(protectedRouter)
	}
}
