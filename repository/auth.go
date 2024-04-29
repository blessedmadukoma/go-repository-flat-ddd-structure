package repository

import (
	"errors"
	"goRepositoryPattern/database/models"
	database "goRepositoryPattern/database/sqlc"
	"goRepositoryPattern/messages"
	"goRepositoryPattern/util"
	"goRepositoryPattern/validators"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthRepository interface {
	Register(ctx *gin.Context, arg validators.RegisterInput) (models.RegisterResponse, error)
	Login(ctx *gin.Context, arg validators.LoginInput) (models.LoginResponse, error)
	ResendRegistrationOTP(ctx *gin.Context, arg validators.ResendRegistrationOtpInput) (models.ResendRegistrationOtpResponse, error)
}

func (r *Repository) Register(ctx *gin.Context, arg validators.RegisterInput) (models.RegisterResponse, error) {

	// check if user account exists (I think I should do it here - handle the business logic or repository - this is only supposed to create account)
	dbUser, _ := r.DB.GetAccountByEmail(ctx, arg.Email)

	if dbUser.ID != 0 {
		return models.RegisterResponse{}, messages.ErrUserExists
	}

	// hash password
	hashedPassword, err := util.HashPassword(arg.Password)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	args := database.CreateAccountParams{
		Firstname:      arg.FirstName,
		Lastname:       arg.LastName,
		Email:          arg.Email,
		HashedPassword: hashedPassword,
	}

	user, err := r.DB.CreateAccount(ctx, args)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	otp := util.RandomOTP()

	// store OTP
	_, err = r.DB.CreateOtp(ctx, database.CreateOtpParams{
		AccountID: user.ID,
		Otp:       otp,
		Type:      int64(messages.AccountTokenTypeEmailConfirmationKey),
	})
	if err != nil {
		log.Println("unable to create OTP in db:", err)
		return models.RegisterResponse{}, err
	}

	response := models.RegisterResponse{
		ID:        user.ID,
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Email:     user.Email,
		OTP:       otp,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (r *Repository) Login(ctx *gin.Context, arg validators.LoginInput) (models.LoginResponse, error) {
	user, err := r.DB.GetAccountByEmail(ctx, arg.Email)
	if err != nil {
		return models.LoginResponse{}, messages.ErrUserNotExists
	}

	// check if user account exists
	if user.ID < 1 {
		return models.LoginResponse{}, messages.ErrInvalidCredentials
	}

	// check hashed password
	err = util.CheckPassword(arg.Password, user.HashedPassword)
	if err != nil {
		return models.LoginResponse{}, messages.ErrInvalidPassword
	}

	token, err := r.Token.CreateToken(user.ID, time.Minute*15)
	if err != nil {
		return models.LoginResponse{}, err
	}

	response := models.LoginResponse{
		ID:        user.ID,
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Email:     user.Email,
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (r *Repository) ResendRegistrationOTP(ctx *gin.Context, arg validators.ResendRegistrationOtpInput) (models.ResendRegistrationOtpResponse, error) {
	a, err := r.DB.GetAccountByEmail(ctx, arg.Email)
	if err != nil {
		log.Println("error getting account by email:", err)
		return models.ResendRegistrationOtpResponse{}, err
	}

	// GetOTPByTypeAndID
	args := database.GetOtpByAccountIDAndTypeParams{
		AccountID: a.ID,
		Type:      int64(messages.AccountTokenTypeEmailConfirmationKey),
	}

	// Check if email is verified
	_, err = r.DB.GetOtpByAccountIDAndType(ctx, args)
	if err != nil {
		// check if err is not found
		if errors.Is(err, messages.ErrRecordNotFound) {
			log.Println("email is already verified, no record in db:", err)
			return models.ResendRegistrationOtpResponse{}, messages.ErrEmailIsVerified
		}
		log.Println("no record in db:", err)
		return models.ResendRegistrationOtpResponse{}, err
	}

	otp := util.RandomOTP()

	newOTPdata := database.UpdateOtpParams{
		Otp:       otp,
		AccountID: a.ID,
		Type:      int64(messages.AccountTokenTypeEmailConfirmationKey),
	}

	_, err = r.DB.UpdateOtp(ctx, newOTPdata)
	if err != nil {
		log.Println("unable to update OTP table:", err)
		return models.ResendRegistrationOtpResponse{}, err
	}

	response := models.ResendRegistrationOtpResponse{
		FirstName: a.Firstname,
		Email:     a.Email,
		OTP:       otp,
	}

	return response, nil
}
