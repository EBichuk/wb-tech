package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HttpServer
}

type HttpServer struct {
	Addr         string        `env:"HTTP_ADDRESS"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT"`
	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT"`
}

func LoadConfig() (Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
