package admin

import (
	"goRepositoryPattern/repository"
	"goRepositoryPattern/services"
)

type AdminController struct {
	repo     repository.Repository
	services services.Service
}

func NewAdminController(repo *repository.Repository, services *services.Service) *AdminController {
	return &AdminController{
		repo:     *repo,
		services: *services,
	}
}
