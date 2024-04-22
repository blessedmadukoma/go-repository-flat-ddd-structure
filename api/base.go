package api

import (
	"goRepositoryPattern/repository"
	"goRepositoryPattern/services"
)

type Controller struct {
	repo     repository.Repository
	services services.Service
}

func NewController(repo *repository.Repository, services *services.Service) *Controller {
	return &Controller{
		repo:     *repo,
		services: *services,
	}
}
