package main

import (
	"context"
	"log"
	"social_network_for_programmers"
	"social_network_for_programmers/internal/config"
	v1 "social_network_for_programmers/internal/delivery/http/v1"
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

	ctx := context.Background()
	pg := postgres.NewPostgres(ctx, cfg.PG)
	client, err := pg.Connection()
	if err != nil {
		log.Fatal(err.Error())
	}

	//if err := pg.MigrationUp(); err != nil && err.Error() != "no changes" {
	//	log.Fatalf("Failed migrations up: %s", err.Error())
	//}

	tokenManager, err := auth.NewManager(cfg.SecretKey)
	if err != nil {
		log.Fatalf("failed to create tokenManager: %s", err.Error())
	}

	repos := repository.NewRepositories(client)
	services := service.NewServices(repos, tokenManager)
	handler := v1.NewHandler(services)

	srv := new(social_network_for_programmers.Server)
	if err = srv.Run(cfg.HttpServer, handler.InitRoutes()); err != nil {
		log.Fatalf("failed to run http server: %s", err.Error())
	}
}
