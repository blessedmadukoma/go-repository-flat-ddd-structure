package models

import (
	database "goRepositoryPattern/database/sqlc"
	"time"
)

type RegisterResponse struct {
	ID        int64         `json:"-"`
	FirstName string        `json:"firstname"`
	LastName  string        `json:"lastname"`
	Email     string        `json:"email"`
	Role      database.Role `json:"-"`
	OTP       string        `json:"-"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type LoginResponse struct {
	ID        int64         `json:"-"`
	FirstName string        `json:"firstname"`
	LastName  string        `json:"lastname"`
	Email     string        `json:"email"`
	Role      database.Role `json:"-"`
	Token     string        `json:"token"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type ResendRegistrationOtpResponse struct {
	FirstName string    `json:"firstname"`
	Email     string    `json:"email"`
	OTP       string    `json:"otp"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VerifyAccountResponse struct {
	FirstName string    `json:"firstname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ForgotPasswordResponse struct {
	FirstName string `json:"firstname"`
	Email     string `json:"email"`
	Link      string `json:"reset_link"`
}

type ForgotPasswordConfirmResponse struct {
	Message string `json:"message"`
}

type ForgotPasswordChangeResponse struct {
	Message string `json:"message"`
}
