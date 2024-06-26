package messages

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	OperationWasSuccessful = "operation was successful"
	NotFound               = "not found"
	SomethingWentWrong     = "something went wrong, please contact support"
	InvalidCredentials     = "invalid credentials"
	EmailIsVerified        = "email is verified, please proceed to login"
	UnAuthorizedAccess     = "unauthorized access"

	ErrEmailIsVerified    = errors.New("email is verified, please proceed to login")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidOTP         = errors.New("invalid account OTP")
	ErrInvalidToken       = errors.New("invalid authentication token")
	ErrInvalidLink        = errors.New("invalid link")
	ErrInvalidPassword    = errors.New("password does not match")
	ErrExpiredToken       = errors.New("token expired")
	ErrValidationFailed   = errors.New("validation failed, please contact support")
	ErrUserExists         = errors.New("user already exists")
	ErrUserNotExists      = errors.New("user does not exist")
)

const (
	ForeignKeyViolation = "23503"
	UniqueViolation     = "23505"
)

var ErrRecordNotFound = pgx.ErrNoRows
var ErrUniqueViolation = &pgconn.PgError{
	Code: UniqueViolation,
}

func ErrorCode(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code
	}

	return ""
}
