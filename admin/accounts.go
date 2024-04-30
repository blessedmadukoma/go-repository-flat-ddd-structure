package admin

import (
	"goRepositoryPattern/messages"
	"goRepositoryPattern/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ac AdminController) GetAccounts(ctx *gin.Context) {
	var R = messages.ResponseFormat{}

	var req validators.ListAccountInput
	if err := ctx.ShouldBindJSON(&req); err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ErrValidationFailed.Error()
		ctx.JSON(messages.Response(http.StatusUnprocessableEntity, R))
		return
	}

	response, err := ac.services.AccountService.GetAccounts(ctx, req)
	if err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ErrValidationFailed.Error()
		ctx.JSON(messages.Response(http.StatusUnprocessableEntity, R))
		return
	}

	R.Data = response

	ctx.JSON(messages.Response(http.StatusOK, R))
}
