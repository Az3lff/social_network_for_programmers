package main

import (
	"social_network_for_programmers"
	"social_network_for_programmers/internal/Messanger"
	"social_network_for_programmers/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {

	}

	router := gin.New()
	Messagehandler := Messanger.NewMessangerHandler()
	Messagehandler.Register(router)
	srv := new(social_network_for_programmers.Server)
	if err = srv.Run(cfg.HttpServer, router); err != nil {

	}
}
