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
		R.Message = messages.ValidationFailed
		ctx.JSON(messages.Response(http.StatusUnprocessableEntity, R))
		return
	}

	user, err := c.services.AuthService.Register(ctx, registerArg)
	if err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = err.Error()
		ctx.JSON(messages.Response(http.StatusUnauthorized, R))
		return
	}

	R.Data = user

	ctx.JSON(messages.Response(http.StatusCreated, R))
}

func (c *Controller) Login(ctx *gin.Context) {
	var R = messages.ResponseFormat{}

	// Parse request parameters
	var loginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.BindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service method to handle login
	authToken, err := c.services.AuthService.Login(ctx, loginRequest.Email, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// If login is successful, return authentication token
	R.Data = gin.H{
		"token": authToken,
	}

	ctx.JSON(messages.Response(0, R))
}
