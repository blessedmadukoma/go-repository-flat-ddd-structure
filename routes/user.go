package routes

import "github.com/gin-gonic/gin"

func registerUserRoute(rg *gin.RouterGroup) {

	router := rg.Group("/users")

	router.POST("", m.Throttle(2), c.GetUserByID)
}
