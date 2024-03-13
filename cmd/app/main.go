package main

import (
	"social_network_for_programmers"
	"social_network_for_programmers/internal/Messanger"
	"social_network_for_programmers/internal/config"
	"social_network_for_programmers/internal/service/authentication/models"
	"social_network_for_programmers/internal/service/authentication"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {

	}

	router := gin.New()
	messageHandler := Messanger.NewMessangerHandler()
	messageHandler.Register(router)
	authHandler := &authentication.AuthHandler{&authentication.AuthStorage{map[int]models.User{}}}
	router.RouterGroup.POST("/register", authHandler.Register)
	srv := new(social_network_for_programmers.Server)
	if err = srv.Run(cfg.HttpServer, router); err != nil {

	}
}
