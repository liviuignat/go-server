package controllers

import (
	. "go-server/src/models"

	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
)

func GetUserHandler(rendr render.Render, db *gorm.DB) {
	users := []User{
		User{FirstName: "Lau", LastName: "Bau"},
		User{FirstName: "Alina", LastName: "Bau"},
		User{FirstName: "Liviu", LastName: "Varza"},
	}

	db.DropTable(&User{})
	db.CreateTable(&User{})
	for _, user := range users {
		db.Create(&user)
	}

	firstUser := User{}
	db.First(&firstUser)

	rendr.JSON(200, firstUser)
}
