package main

import "github.com/caarlos0/env"

type Config struct {
	Port               string `env:"PORT" envDefault:"3000"`
	DatabaseConnection string `env:"DB_CONNECTION" envDefault:"user=postgres password=password dbname=gorm sslmode=disable port=5433"`
}

func GetConfig() Config {
	cfg := Config{}
	env.Parse(&cfg)
	return cfg
}
