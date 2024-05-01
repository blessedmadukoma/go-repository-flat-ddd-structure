package routes

import (
	"github.com/gin-gonic/gin"
)

func registerAdminRoute(rg *gin.RouterGroup) {

	router := rg.Group("/admin")
	{
		// admin/auth
		auth := router.Group("/auth")
		auth.Use(m.Throttle(5))
		{
			auth.POST("/register", c.Register)
			auth.POST("/login", ac.Login)
		}

		// admin/accounts
		accounts := router.Group("/accounts")
		accounts.Use(m.AdminMiddleware())
		// accounts.Use(m.AuthMiddleware())
		{

			accounts.GET("", ac.GetAccounts)
			// accounts.PUT("/approve/:id", ac.ApproveAccount)
			accounts.GET("/:id", ac.GetAccountByID)
		}
	}
}
