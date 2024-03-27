package v1

import (
	"social_network_for_programmers/internal/config"
	"social_network_for_programmers/internal/middleware/auth"
	"social_network_for_programmers/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
	cfg      *config.Config
}

func NewHandler(services *service.Services, cfg *config.Config) *Handler {
	return &Handler{services, cfg}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initAuthRoutes(v1)
		protectedRouter := v1.Group("")
		protectedRouter.Use(auth.JwtAuthMiddleware(h.cfg.SecretKey))

		h.initMessengerRoutes(protectedRouter)
	}
}
