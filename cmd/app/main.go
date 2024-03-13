package main

import (
	"social_network_for_programmers"
	"social_network_for_programmers/internal/config"
	v1 "social_network_for_programmers/internal/delivery/http/v1"
	"social_network_for_programmers/internal/repository"
	"social_network_for_programmers/internal/service"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {

	}

	repos := repository.NewRepositories()
	services := service.NewServices(repos)
	handlers := v1.NewHandler(services)

	srv := new(social_network_for_programmers.Server)
	if err = srv.Run(cfg.HttpServer, handlers.InitRoutes()); err != nil {

	}
}
