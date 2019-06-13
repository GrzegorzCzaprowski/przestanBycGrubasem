package main

import (
	handlers "github.com/GrzegorzCzaprowski/przestanBycGrubasem/handlers/users"
	"github.com/GrzegorzCzaprowski/przestanBycGrubasem/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	var model models.Model = 3

	userHandler := handlers.UserHandler{M: model}

	router.POST("/user/post", userHandler.Post)
	router.POST("/user/login", userHandler.Log)
}
