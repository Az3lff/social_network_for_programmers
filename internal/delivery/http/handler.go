package http

import (
	"github.com/gin-gonic/gin"
	"social_network_for_programmers/internal/config"
	v1 "social_network_for_programmers/internal/delivery/http/v1"
	"social_network_for_programmers/internal/middleware/auth"
	"social_network_for_programmers/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

// http://localhost:8080/api/v1/...
func (h *Handler) InitRoutes(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		auth.AuthMiddleware(cfg.SecretKey),
	)

	handlerV1 := v1.NewHandler(h.services)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}

	return router
}
