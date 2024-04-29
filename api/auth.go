package api

import (
	"goRepositoryPattern/messages"
	"goRepositoryPattern/tasks"
	"goRepositoryPattern/util"
	"goRepositoryPattern/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
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

	// Send email
	go func() {
		tasks.RegisterOtpTask(tasks.RegisterOtpInput{
			Email:     user.Email,
			FirstName: user.FirstName,
			OTP:       user.OTP,
		})
	}()

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

	user, err := c.services.AuthService.Login(ctx, loginArg)
	if err != nil {
		R.Error = append(R.Error, err.Error())
		// R.Message = err.Error()
		R.Message = messages.SomethingWentWrong
		ctx.JSON(messages.Response(http.StatusUnauthorized, R))
		return
	}

	// Resend registration token
	go func() {
		tasks.RegisterOtpTask(tasks.RegisterOtpInput{
			Email:     user.Email,
			FirstName: user.FirstName,
			OTP:       util.RandomOTP(),
		})
	}()

	R.Data = user

	ctx.JSON(messages.Response(0, R))
}

func (c Controller) ResendRegistrationOtp(ctx *gin.Context) {
	var R = messages.ResponseFormat{}

	// Validate input
	var i validators.ResendRegistrationOtpInput
	if err := ctx.ShouldBindJSON(&i); err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ErrValidationFailed.Error()
		ctx.JSON(messages.Response(http.StatusUnprocessableEntity, R))
		return
	}

	user, err := c.services.AuthService.ResendRegistrationOtp(ctx, i)
	if err != nil {
		switch err {
		case messages.ErrEmailIsVerified:
			R.Message = messages.EmailIsVerified
			ctx.JSON(messages.Response(http.StatusBadRequest, R))
			return
		case pgx.ErrNoRows:
			R.Message = messages.ErrRecordNotFound.Error()
			ctx.JSON(messages.Response(http.StatusBadRequest, R))
			return
		default:
			R.Message = messages.EmailIsVerified
			ctx.JSON(messages.Response(http.StatusOK, R))
			return
		}

	}

	// Resend registration token
	go func() {
		tasks.RegisterOtpTask(tasks.RegisterOtpInput{
			Email:     user.Email,
			FirstName: user.FirstName,
			OTP:       user.OTP,
		})
	}()

	ctx.JSON(messages.Response(http.StatusOK, R))
}

func (c Controller) VerifyAccount(ctx *gin.Context) {
	var R = messages.ResponseFormat{}

	// Validate input
	var i validators.VerifyAccountInput
	if err := ctx.ShouldBindJSON(&i); err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ErrValidationFailed.Error()
		ctx.JSON(messages.Response(http.StatusUnprocessableEntity, R))
		return
	}

	user, err := c.services.AuthService.VerifyAccount(ctx, i)
	if err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.SomethingWentWrong
		ctx.JSON(messages.Response(http.StatusUnauthorized, R))
		return
	}

	R.Data = user

	ctx.JSON(messages.Response(http.StatusOK, R))
}
