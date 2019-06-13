package handlers

import (
	"github.com/GrzegorzCzaprowski/przestanBycGrubasem/models"
)

type modelerWeights interface {
	AddWeight(models.User, models.Weight) error
}

type WeightHandler struct {
	M modelerWeights
}
