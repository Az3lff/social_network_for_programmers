package main

import (
	"context"
	"fmt"
	"log"
	"social_network_for_programmers"
	"social_network_for_programmers/internal/config"
	"social_network_for_programmers/internal/delivery/http"
	"social_network_for_programmers/internal/repository"
	"social_network_for_programmers/internal/service"
	"social_network_for_programmers/pkg/auth"
	"social_network_for_programmers/pkg/database/postgres"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to get config: %s", err.Error())
	}
	fmt.Println(cfg)

	ctx := context.Background()
	pg := postgres.NewPostgres(ctx, &cfg.PG)
	client, err := pg.Connection()
	if err != nil {
		log.Fatal(err.Error())
	}

	tokenManager, err := auth.NewManager(cfg.SecretKey)
	if err != nil {
		log.Fatalf("failed to create tokenManager: %s", err.Error())
	}

	repos := repository.NewRepositories(client)
	services := service.NewServices(repos, tokenManager)
	handler := http.NewHandler(services)

	srv := new(social_network_for_programmers.Server)
	if err = srv.Run(cfg.HttpServer, handler.InitRoutes(cfg)); err != nil {
		log.Fatalf("failed to run http server: %s", err.Error())
	}
}
