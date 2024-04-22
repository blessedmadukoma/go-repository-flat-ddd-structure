package middleware

import (
	"goRepositoryPattern/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m Middleware) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if origin := c.Request.Header.Get("Origin"); origin != "" {
			originMap := make(map[string]string)
			origins := strings.Split(util.CorsWhiteList(), ",")

			for _, s := range origins {
				originMap[s] = s
			}
			c.Writer.Header().Set("Access-Control-Allow-Origin", originMap[origin])
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Range")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
