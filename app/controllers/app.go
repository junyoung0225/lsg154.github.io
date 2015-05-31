package controllers

import (
	"lsg154.github.io/app/models"
)

type App struct {
	GormController
	CurrentUser *models.User
}

