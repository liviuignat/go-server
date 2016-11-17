package main

import (
	"go-server/src/controllers"
	"log"

	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/render"
)

func main() {
	app := martini.Classic()
	app.Use(render.Renderer())
	app.Use(databaseMiddleware())

	app.Get("/", controllers.GetHomeHandler)
	app.Group("/users", func(router martini.Router) {
		router.Get("", controllers.GetUserHandler)
	})

	app.Run()
}

func databaseMiddleware() martini.Handler {
	return func(context martini.Context) {
		db, err := gorm.Open("postgres", GetConfig().DatabaseConnection)

		if err != nil {
			log.Fatalln(err.Error())
		}

		defer db.Close()
		db.LogMode(true)

		context.Map(db)

		context.Next()
	}
}
