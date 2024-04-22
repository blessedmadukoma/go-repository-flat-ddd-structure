package middleware

import (
	"goRepositoryPattern/api"
	"goRepositoryPattern/repository"
)

type Middleware struct {
	repo *repository.Repository
	c    api.Controller
}

func NewMiddleware(repo *repository.Repository, c api.Controller) Middleware {
	return Middleware{repo, c}
}
