package repository

import (
	"errors"
	"fmt"
	"goRepositoryPattern/database/models"
	database "goRepositoryPattern/database/sqlc"
	"goRepositoryPattern/messages"
	"goRepositoryPattern/util"
	"goRepositoryPattern/validators"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthRepository interface {
	Register(ctx *gin.Context, arg validators.RegisterInput) (models.RegisterResponse, error)
	Login(ctx *gin.Context, arg validators.LoginInput) (models.LoginResponse, error)
	ResendRegistrationOTP(ctx *gin.Context, arg validators.ResendRegistrationOtpInput) (models.ResendRegistrationOtpResponse, error)
	VerifyAccount(ctx *gin.Context, arg validators.VerifyAccountInput) (models.VerifyAccountResponse, error)
	PasswordReset(ctx *gin.Context, arg validators.PasswordResetInput) (models.ForgotPasswordResponse, error)
	PasswordResetConfirm(ctx *gin.Context, arg validators.PasswordResetConfirmInput) (models.ForgotPasswordConfirmResponse, error)
	PasswordResetChange(ctx *gin.Context, arg validators.PasswordResetChangeInput) (models.ForgotPasswordChangeResponse, error)

	// Admin routes
	AdminLogin(ctx *gin.Context, arg validators.AdminLoginInput) (models.LoginResponse, error)
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

	var role database.Role

	if arg.Role == database.RoleAdmin {
		role = database.RoleAdmin
	} else {
		role = database.RoleUser
	}

	args := database.CreateAccountParams{
		Firstname:      arg.FirstName,
		Lastname:       arg.LastName,
		Email:          arg.Email,
		Role:           role,
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
		Role:      user.Role,
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
		Role:      user.Role,
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (r *Repository) ResendRegistrationOTP(ctx *gin.Context, arg validators.ResendRegistrationOtpInput) (models.ResendRegistrationOtpResponse, error) {
	a, err := r.DB.GetAccountByEmail(ctx, arg.Email)
	if err != nil {
		if errors.Is(err, messages.ErrRecordNotFound) {
			return models.ResendRegistrationOtpResponse{}, messages.ErrUserNotExists
		}

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
			return models.ResendRegistrationOtpResponse{}, messages.ErrEmailIsVerified
		}

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
		return models.ResendRegistrationOtpResponse{}, err
	}

	response := models.ResendRegistrationOtpResponse{
		FirstName: a.Firstname,
		Email:     a.Email,
		OTP:       otp,
	}

	return response, nil
}

func (r *Repository) VerifyAccount(ctx *gin.Context, arg validators.VerifyAccountInput) (models.VerifyAccountResponse, error) {
	a, err := r.DB.GetAccountByEmail(ctx, arg.Email)
	if err != nil {
		if errors.Is(err, messages.ErrRecordNotFound) {
			return models.VerifyAccountResponse{}, messages.ErrUserNotExists
		}

		return models.VerifyAccountResponse{}, err
	}

	// GetOTPByTypeAndID
	args := database.GetOtpByAccountIDAndTypeParams{
		AccountID: a.ID,
		Type:      int64(messages.AccountTokenTypeEmailConfirmationKey),
	}

	// Check if email is verified
	otp, err := r.DB.GetOtpByAccountIDAndType(ctx, args)
	if err != nil {
		// check if err is not found
		if errors.Is(err, messages.ErrRecordNotFound) {
			return models.VerifyAccountResponse{}, messages.ErrEmailIsVerified
		}

		return models.VerifyAccountResponse{}, err
	}

	if otp.Otp != arg.OTP {
		return models.VerifyAccountResponse{}, messages.ErrInvalidOTP
	}

	// update account status
	_, err = r.DB.UpdateAccountStatus(ctx, database.UpdateAccountStatusParams{
		ID: a.ID,
		IsVerified: pgtype.Bool{
			Bool:  true,
			Valid: true,
		},
	})
	if err != nil {
		return models.VerifyAccountResponse{}, err
	}

	// delete OTP
	err = r.DB.DeleteOtp(ctx, database.DeleteOtpParams{
		ID:        otp.ID,
		AccountID: otp.AccountID,
		Type:      int64(messages.AccountTokenTypeEmailConfirmationKey),
	})
	if err != nil {
		return models.VerifyAccountResponse{}, err
	}

	response := models.VerifyAccountResponse{
		FirstName: a.Firstname,
		Email:     a.Email,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}

	return response, nil
}

func (r Repository) PasswordReset(ctx *gin.Context, arg validators.PasswordResetInput) (models.ForgotPasswordResponse, error) {
	// get user details from db
	a, err := r.DB.GetAccountByEmail(ctx, arg.Email)
	if err != nil {
		if errors.Is(err, messages.ErrRecordNotFound) {
			return models.ForgotPasswordResponse{}, messages.ErrUserNotExists
		}

		return models.ForgotPasswordResponse{}, err
	}

	// Get OTP/Reset Link By Type And ID
	args := database.GetOtpByAccountIDAndTypeParams{
		AccountID: a.ID,
		Type:      int64(messages.AccountTokenTypePasswordResetKey),
	}

	// get link if any (if there is any, generate a new link and send)
	dbLink, err := r.DB.GetOtpByAccountIDAndType(ctx, args)
	if err != nil {
		log.Println("reset link not found, proceeding to generate one:", err)
	}

	// generate token
	link := util.RandomResetLink()

	// store in db
	if dbLink.Otp != "" {
		// update existing OTP
		_, err = r.DB.UpdateOtp(ctx, database.UpdateOtpParams{
			AccountID: a.ID,
			Otp:       link,
			Type:      int64(messages.AccountTokenTypePasswordResetKey),
		})
		if err != nil {
			log.Println("error updating reset link - ", err)
			return models.ForgotPasswordResponse{}, err
		}
	} else {
		// create new OTP
		_, err = r.DB.CreateOtp(ctx, database.CreateOtpParams{
			AccountID: a.ID,
			Otp:       link,
			Type:      int64(messages.AccountTokenTypePasswordResetKey),
		})
		if err != nil {
			log.Println("error in creating reset link - ", err)
			return models.ForgotPasswordResponse{}, err
		}
	}

	response := models.ForgotPasswordResponse{
		FirstName: a.Firstname,
		Email:     a.Email,
		Link:      fmt.Sprintf("%s/%s", os.Getenv("RESET_LINK"), link),
	}

	return response, nil

}

// PasswordResetConfirm confirms the password reset token
func (r Repository) PasswordResetConfirm(ctx *gin.Context, arg validators.PasswordResetConfirmInput) (models.ForgotPasswordConfirmResponse, error) {
	// get the user by email
	a, err := r.DB.GetAccountByEmail(ctx, arg.Email)
	if err != nil {
		if errors.Is(err, messages.ErrRecordNotFound) {
			return models.ForgotPasswordConfirmResponse{}, messages.ErrUserNotExists
		}

		return models.ForgotPasswordConfirmResponse{}, err
	}

	// Get Reset Link By Type And ID
	args := database.GetOtpByAccountIDAndTypeParams{
		AccountID: a.ID,
		Type:      int64(messages.AccountTokenTypePasswordResetKey),
	}

	linkData, err := r.DB.GetOtpByAccountIDAndType(ctx, args)
	if err != nil {
		return models.ForgotPasswordConfirmResponse{}, messages.ErrInvalidLink
	}

	if linkData.Otp != arg.Link {
		return models.ForgotPasswordConfirmResponse{}, messages.ErrInvalidLink
	}

	response := models.ForgotPasswordConfirmResponse{
		Message: "Link is valid",
	}

	return response, nil
}

// PasswordResetChange changes the password using the reset link
func (r Repository) PasswordResetChange(ctx *gin.Context, arg validators.PasswordResetChangeInput) (models.ForgotPasswordChangeResponse, error) {
	// get the user by email
	a, err := r.DB.GetAccountByEmail(ctx, arg.Email)
	if err != nil {
		if errors.Is(err, messages.ErrRecordNotFound) {
			return models.ForgotPasswordChangeResponse{}, messages.ErrUserNotExists
		}

		return models.ForgotPasswordChangeResponse{}, err
	}

	// Get Reset Link By Type And ID
	args := database.GetOtpByAccountIDAndTypeParams{
		AccountID: a.ID,
		Type:      int64(messages.AccountTokenTypePasswordResetKey),
	}

	linkData, err := r.DB.GetOtpByAccountIDAndType(ctx, args)
	if err != nil {
		return models.ForgotPasswordChangeResponse{}, messages.ErrInvalidLink
	}

	if linkData.Otp != arg.Link {
		return models.ForgotPasswordChangeResponse{}, messages.ErrInvalidLink
	}

	// hash password
	hashedPassword, err := util.HashPassword(arg.NewPassword)
	if err != nil {
		return models.ForgotPasswordChangeResponse{}, err
	}

	// update password
	_, err = r.DB.UpdateAccountPassword(ctx, database.UpdateAccountPasswordParams{
		ID:             a.ID,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		return models.ForgotPasswordChangeResponse{}, err
	}

	// delete OTP
	err = r.DB.DeleteOtp(ctx, database.DeleteOtpParams{
		ID:        linkData.ID,
		AccountID: linkData.AccountID,
		Type:      int64(messages.AccountTokenTypePasswordResetKey),
	})
	if err != nil {
		return models.ForgotPasswordChangeResponse{}, err
	}

	response := models.ForgotPasswordChangeResponse{
		Message: "Password changed successfully",
	}

	return response, nil
}
