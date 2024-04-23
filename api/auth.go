package api

import (
	"goRepositoryPattern/messages"
	"goRepositoryPattern/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Register(ctx *gin.Context) {
	var R = messages.ResponseFormat{}

	var registerArg validators.RegisterInput

	if err := ctx.ShouldBindJSON(&registerArg); err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ErrValidationFailed.Error()
		ctx.JSON(messages.Response(http.StatusUnprocessableEntity, R))
		return
	}

	user, err := c.services.AuthService.Register(ctx, registerArg)
	if err != nil {
		R.Error = append(R.Error, err.Error())
		// R.Message = err.Error()
		R.Message = messages.SomethingWentWrong
		ctx.JSON(messages.Response(http.StatusUnauthorized, R))
		return
	}

	R.Data = user

	ctx.JSON(messages.Response(http.StatusCreated, R))
}

func (c *Controller) Login(ctx *gin.Context) {
	var R = messages.ResponseFormat{}

	var loginArg validators.LoginInput

	if err := ctx.ShouldBindJSON(&loginArg); err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ErrValidationFailed.Error()
		ctx.JSON(messages.Response(http.StatusUnprocessableEntity, R))
		return
	}

	// Call service method to handle login
	user, err := c.services.AuthService.Login(ctx, loginArg)
	if err != nil {
		R.Error = append(R.Error, err.Error())
		// R.Message = err.Error()
		R.Message = messages.SomethingWentWrong
		ctx.JSON(messages.Response(http.StatusUnauthorized, R))
		return
	}

	// If login is successful, return authentication token
	R.Data = user

	ctx.JSON(messages.Response(0, R))
}
