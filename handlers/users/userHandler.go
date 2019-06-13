package handlers

import (
	"github.com/GrzegorzCzaprowski/przestanBycGrubasem/models"
)

type modelerUsers interface {
	PostUser(models.User) (models.User, error)
	LogUser(models.User) (models.User, error)
}
type UserHandler struct {
	M modelerUsers
}
