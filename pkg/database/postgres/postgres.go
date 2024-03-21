package postgres

import (
	"context"
	"errors"
	"fmt"
	"social_network_for_programmers/internal/config"
	// "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	client *pgxpool.Pool
	pgCfg  config.PG
	ctx    context.Context
	url    string
}

func NewPostgres(ctx context.Context, cfg config.PG) *Postgres {
	return &Postgres{
		ctx:   ctx,
		pgCfg: cfg,
	}
}

func (p *Postgres) Connection() (*pgxpool.Pool, error) {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", p.pgCfg.Username, p.pgCfg.Password, p.pgCfg.Host, p.pgCfg.Port, p.pgCfg.DbName)
	pgxConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, errors.New("failed to parse postgres config: " + err.Error())
	}

	client, err := pgxpool.NewWithConfig(p.ctx, pgxConfig)
	if err != nil {
		return nil, errors.New("failed to connect to the postgres: " + err.Error())
	}

	p.client = client
	p.url = url

	return p.client, nil
}

func (p *Postgres) MigrationUp() error {
	return nil
}

func (p *Postgres) MigrationDown() error {
	return nil
}
