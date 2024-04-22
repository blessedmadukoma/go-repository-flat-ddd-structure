package messages

import "errors"

var (
	OperationWasSuccessful = "operation was successful"
	NotFound               = "not found"
	SomethingWentWrong     = "something went wrong, please contact support"
	InvalidCredentials     = "invalid credentials"
	ValidationFailed       = "validation failed, please contact support"

	ErrInvalidToken = errors.New("invalid authentication token")
	ErrExpiredToken = errors.New("token expired")
)
