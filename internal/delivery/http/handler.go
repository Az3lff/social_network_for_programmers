package http

import (
	"github.com/gin-gonic/gin"
	"social_network_for_programmers/internal/config"
	v1 "social_network_for_programmers/internal/delivery/http/v1"
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

// http://localhost:8080/snp/v1/...

func (h *Handler) InitRoutes(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.NewHandler(h.services, cfg)
	api := router.Group("/snp")
	{
		handlerV1.Init(api)
	}

	return router
}
