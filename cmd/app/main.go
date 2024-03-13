package main

import (
	"github.com/gin-gonic/gin"
	"social_network_for_programmers"
	"social_network_for_programmers/internal/config"
	"social_network_for_programmers/internal/service/authentication"
	"social_network_for_programmers/internal/service/authentication/models"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {

	}

	router := gin.New()

	authHandler := &authentication.AuthHandler{&authentication.AuthStorage{map[int]models.User{}}}
	router.RouterGroup.POST("/register", authHandler.Register)

	srv := new(social_network_for_programmers.Server)
	if err = srv.Run(cfg.HttpServer, router); err != nil {

	}
}
