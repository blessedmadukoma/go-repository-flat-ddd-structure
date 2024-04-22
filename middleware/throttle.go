package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joshuachinemezu/throttle"
)

func (m Middleware) Throttle(limit uint64) gin.HandlerFunc {
	return throttle.Policy(&throttle.Quota{
		Limit:  limit,
		Within: time.Minute,
	})
}
