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

	ErrInvalidToken     = errors.New("invalid authentication token")
	ErrExpiredToken     = errors.New("token expired")
	ErrValidationFailed = errors.New("validation failed, please contact support")
	ErrUserExists       = errors.New("user already exists")
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
