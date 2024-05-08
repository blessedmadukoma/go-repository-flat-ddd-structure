package admin

import (
	"goRepositoryPattern/messages"
	"goRepositoryPattern/validators"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ac AdminController) GetAccounts(ctx *gin.Context) {
	var R = messages.ResponseFormat{}

	var req validators.ListAccountInput
	if err := ctx.ShouldBindQuery(&req); err != nil {
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

func (ac AdminController) GetAccountByID(ctx *gin.Context) {
	var R = messages.ResponseFormat{}

	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ErrValidationFailed.Error()
		ctx.JSON(messages.Response(http.StatusUnprocessableEntity, R))
		return
	}

	response, err := ac.services.AccountService.GetAccountByID(ctx, int64(intID))
	if err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ErrValidationFailed.Error()
		ctx.JSON(messages.Response(http.StatusUnprocessableEntity, R))
		return
	}

	R.Data = response

	ctx.JSON(messages.Response(http.StatusOK, R))
}

func (ac AdminController) GetAccountByEmail(ctx *gin.Context) {
	var R = messages.ResponseFormat{}

	type getAccountByEmailInput struct {
		Email string `json:"email" binding:"required"`
	}

	var req getAccountByEmailInput
	if err := ctx.ShouldBindJSON(&req); err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ErrValidationFailed.Error()
		ctx.JSON(messages.Response(http.StatusUnprocessableEntity, R))
		return
	}

	response, err := ac.services.AccountService.GetAccountByEmail(ctx, req.Email)
	if err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ErrValidationFailed.Error()
		ctx.JSON(messages.Response(http.StatusUnprocessableEntity, R))
		return
	}

	R.Data = response

	ctx.JSON(messages.Response(http.StatusOK, R))
}

func (ac AdminController) DeleteAccount(ctx *gin.Context) {
	var R = messages.ResponseFormat{}

	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ErrValidationFailed.Error()
		ctx.JSON(messages.Response(http.StatusUnprocessableEntity, R))
		return
	}

	err = ac.services.AccountService.DeleteAccount(ctx, int64(intID))
	if err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ErrValidationFailed.Error()
		ctx.JSON(messages.Response(http.StatusUnprocessableEntity, R))
		return
	}

	ctx.JSON(messages.Response(http.StatusOK, R))
}
