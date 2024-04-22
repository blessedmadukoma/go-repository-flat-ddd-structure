package routes

import (
	"goRepositoryPattern/api"
	"goRepositoryPattern/middleware"
	"goRepositoryPattern/repository"
	"goRepositoryPattern/services"

	"github.com/gin-gonic/gin"
)

var (
	c api.Controller
	m middleware.Middleware
	// ac admin.AdminController
)

func RegisterRoutes(engine *gin.Engine, repo *repository.Repository, services *services.Service) {
	c = *api.NewController(repo, services)
	m = middleware.NewMiddleware(repo, c)

	engine.Use(m.CORSMiddleware())

	v1 := engine.Group("/v1")

	registerAuthRoute(v1)
	registerUserRoute(v1)
}
