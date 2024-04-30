package admin

import (
	"net/http"

	"goRepositoryPattern/messages"
	"goRepositoryPattern/validators"

	"github.com/gin-gonic/gin"
)

func (c AdminController) Login(ctx *gin.Context) {
	var R = messages.ResponseFormat{}

	var i validators.AdminLoginInput
	if err := ctx.ShouldBindJSON(&i); err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ErrValidationFailed.Error()
		ctx.JSON(messages.Response(http.StatusUnprocessableEntity, R))
		return
	}

	response, err := c.services.AuthService.AdminLogin(ctx, i)
	if err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.SomethingWentWrong
		ctx.JSON(messages.Response(http.StatusBadRequest, R))
		return
	}

	R.Data = response
	ctx.JSON(messages.Response(http.StatusOK, R))

}
