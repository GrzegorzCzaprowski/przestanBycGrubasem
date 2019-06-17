package main

import (
	"log"
	"net/http"

	userHandlers "github.com/GrzegorzCzaprowski/przestanBycGrubasem/handlers/users"
	weightHandlers "github.com/GrzegorzCzaprowski/przestanBycGrubasem/handlers/weights"

	"github.com/GrzegorzCzaprowski/przestanBycGrubasem/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	var baza []models.User
	user1 := models.User{Email: "test1", Password: "test1"}
	model := models.Model{Db: &baza}
	baza = append(baza, user1)

	userHandler := userHandlers.UserHandler{M: model}
	weightHandler := weightHandlers.WeightHandler{M: model}

	router.POST("/user/post", userHandler.Post)
	router.GET("/user/get", userHandler.Get)
	router.POST("/user/login", userHandler.Login)
	router.POST("/user/logout", userHandler.Logout)
	router.POST("/weight/add", weightHandler.Add)

	log.Fatal(http.ListenAndServe(":8000", router))
}
