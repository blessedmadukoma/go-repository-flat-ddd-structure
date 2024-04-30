package routes

import "github.com/gin-gonic/gin"

func registerAuthRoute(rg *gin.RouterGroup) {

	router := rg.Group("/auth")

	router.POST("/register", m.Throttle(4), c.Register)
	router.POST("/login", m.Throttle(4), c.Login)
	router.POST("/resend-registration-otp", m.Throttle(4), c.ResendRegistrationOtp)
	router.POST("/verify-account", m.Throttle(3), c.VerifyAccount)
	router.POST("/password/reset", m.Throttle(2), c.PasswordReset)
	router.POST("/password/reset/confirm", m.Throttle(2), c.PasswordResetConfirm)

	// router.POST("/password/reset/change", m.Throttle(2), c.PasswordResetChange)
	// router.GET("/otp/generate", m.TokenAuthMiddleware(), m.Throttle(2), c.Generate2FAOTP)
	// router.POST("/otp/verify", m.TokenAuthMiddleware(), m.Throttle(4), c.Verify2FAOTP)
	// router.POST("/otp/validate", m.TokenAuthMiddleware(), m.Throttle(4), c.Validate2FAOTP)
	// router.POST("/otp/disable", m.TokenAuthMiddleware(), m.Throttle(2), c.Disable2FAOTP)
}
