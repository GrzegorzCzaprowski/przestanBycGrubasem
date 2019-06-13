package main

import (
	"log"
	"net/http"

	userHandlers "github.com/GrzegorzCzaprowski/przestanBycGrubasem/handlers/users"
	"github.com/GrzegorzCzaprowski/przestanBycGrubasem/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	var baza []models.User

	var model models.Model
	model.Db = baza

	userHandler := userHandlers.UserHandler{M: model}

	router.POST("/user/post", userHandler.Post)
	router.POST("/user/login", userHandler.Login)
	router.POST("/user/logout", userHandler.Logout)
	router.POST("/weight/add", userHandler.Logout)

	log.Fatal(http.ListenAndServe(":8000", router))
}
