package controllers

import (
	"go-server/src/models"

	"github.com/martini-contrib/render"
)

func GetHomeHandler(r render.Render) {
	model := models.Home{Id: 2}

	r.JSON(200, model)
}
