package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	HttpServer `yaml:"http_server"`
	PG         `yaml:"postgres"`
	JwtToken   string `yaml:"jwt_token"`
}

type HttpServer struct {
	Port           string        `yaml:"port"`
	MaxHeaderBytes int           `yaml:"max_header_bytes"`
	ReadTimeout    time.Duration `yaml:"read_timeout"`
	WriteTimeout   time.Duration `yaml:"write_timeout"`
}

type PG struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db_name"`
}

func GetConfig() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig("local.yml", &cfg); err != nil {
		return nil, fmt.Errorf("config could not be read: %s", err.Error())
	}
	return &cfg, nil
}
