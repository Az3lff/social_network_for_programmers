package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"social_network_for_programmers"
	"social_network_for_programmers/internal/config"
	v1 "social_network_for_programmers/internal/delivery/http/v1"
	"social_network_for_programmers/internal/repository"
	"social_network_for_programmers/internal/service"
	"social_network_for_programmers/pkg/auth"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to get config: %s", err.Error())
	}

	connString := fmt.Sprintf("user=%s password=%s dbname=%s", cfg.PG.Username, cfg.PG.Password, cfg.PG.DbName)
	postgresClient, err := pgxpool.New(context.Background(), connString)

	if err != nil {
		//log.Fatalf("failed to connection database: %s", err.Error())
		log.Printf("failed to connection database: %s", err.Error())
	}

	tokenManager, err := auth.NewManager(cfg.JwtToken)
	if err != nil {
		log.Fatalf("failed to create tokenManager: %s", err.Error())
	}

	repos := repository.NewRepositories(postgresClient)
	services := service.NewServices(repos, tokenManager)
	handlers := v1.NewHandler(services)

	srv := new(social_network_for_programmers.Server)
	if err = srv.Run(cfg.HttpServer, handlers.InitRoutes()); err != nil {
		log.Fatalf("failed to run http server: %s", err.Error())
	}
}
