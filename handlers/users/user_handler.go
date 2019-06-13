package handlers

import (
	"github.com/GrzegorzCzaprowski/przestanBycGrubasem/models"
)

type modelerUsers interface {
	PostUser(models.User) (models.User, error)
	LoginUser(models.User) (models.User, error)
	LogoutUser(models.User) error
}
type UserHandler struct {
	M modelerUsers
}
