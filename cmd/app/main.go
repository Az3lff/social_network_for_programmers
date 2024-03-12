package main

import (
	"github.com/gin-gonic/gin"
	"social_network_for_programmers"
	"social_network_for_programmers/internal/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {

	}

	router := gin.New()

	srv := new(social_network_for_programmers.Server)
	if err = srv.Run(cfg.HttpServer, router); err != nil {

	}
}
