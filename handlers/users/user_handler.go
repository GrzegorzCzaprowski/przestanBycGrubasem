package handlers

import (
	"github.com/GrzegorzCzaprowski/przestanBycGrubasem/models"
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("super-secret-key"))

type modelerUsers interface {
	PostUser(models.User) models.User
	LoginUser(models.User) bool
	LogoutUser(models.User) error
	GetAllUsers() []models.User
}
type UserHandler struct {
	M modelerUsers
}
