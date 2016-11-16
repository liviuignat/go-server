package main

import (
	"go-server/src/controllers"

	"github.com/caarlos0/env"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type Config struct {
	Port string `env:"PORT" envDefault:"3000"`
}

func main() {
	app := martini.Classic()
	app.Use(render.Renderer())

	app.Get("/", controllers.HomeController)

	app.Run()
}

func GetConfig() Config {
	cfg := Config{}
	env.Parse(&cfg)
	return cfg
}
