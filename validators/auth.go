package validators

import _ "github.com/go-playground/validator/v10"

// RegisterInput is for validating account creation
type RegisterInput struct {
	FirstName string `json:"firstname" binding:"required" validate:"required,ascii,max=128"`
	LastName  string `json:"lastname" binding:"required" validate:"required,ascii,max=128"`
	Email     string `json:"email" binding:"required" validate:"required,email,max=128"`
	Password  string `json:"password" binding:"required"`
}

// LoginInput is validating account authentication
type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ResendRegistrationOtpInput struct {
	Email string `json:"email" binding:"required"`
}

// VerifyAccountInput for verifying account
type VerifyAccountInput struct {
	Email string `json:"email" binding:"required"`
	OTP   string `json:"otp" binding:"required"`
}

// ConfirmOTPInput for validating OTP confirm
type ConfirmOTPInput struct {
	Phone string `json:"phone" binding:"required"`
	Token string `json:"token" binding:"required"`
}

type PasswordResetInput struct {
	Email string `json:"email" binding:"required"`
}

type PasswordResetConfirmInput struct {
	Email string `json:"email" binding:"required"`
	Link  string `json:"reset_link" binding:"required"`
}

type PasswordResetChangeInput struct {
	Email       string `json:"email" binding:"required"`
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type PasswordUpdateInput struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type Verify2FAOTPInput struct {
	Token string `json:"token" binding:"required"`
}
