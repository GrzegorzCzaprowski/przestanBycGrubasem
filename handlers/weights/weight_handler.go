package handlers

import (
	"github.com/GrzegorzCzaprowski/przestanBycGrubasem/models"
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("super-secret-key"))

type modelerWeights interface {
	AddWeight(models.User, models.Weight)
}

type WeightHandler struct {
	M modelerWeights
}
