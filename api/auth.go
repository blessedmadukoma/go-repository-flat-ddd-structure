package api

import (
	"goRepositoryPattern/messages"
	"goRepositoryPattern/tasks"
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

	// a, err := c.repo.GetAccount("accounts.email = ?", i.Email)
	// if err != nil {
	// 	R.Status = true
	// 	ctx.JSON(c.Response(http.StatusOK, R))
	// 	return
	// }

	// // Check if email is verified
	// _, err = c.repo.GetAccountToken("account_tokens.type = ? AND accounts.id = ?", uint(constants.AccountTokenTypeEmailConfirmationKey), a.ID)
	// if err == gorm.ErrRecordNotFound {
	// 	R.Message = messages.EmailIsVerified
	// 	ctx.JSON(c.Response(http.StatusBadRequest, R))
	// 	return
	// }

	// t := utils.GenerateRandomNumber(4)

	// c.repo.UpdateAccountToken(models.AccountToken{
	// 	Token: t,
	// }, "account_id = ? AND type = ?", a.ID, uint(constants.AccountTokenTypeEmailConfirmationKey))

	// // Resend registration token

}
