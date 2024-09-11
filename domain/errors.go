package domain

import (
	"errors"
)

type CustomError struct {
	Message string
	Err     error
}

func (e CustomError) Error() string {
	return e.Message
}

var (
	ErrNotFound          = errors.New("NotFoundError")
	ErrDuplicateKey      = errors.New("DuplicateKeyError")
	ErrDuplicateUsername = errors.New("DuplicateUsernameError")
	ErrDuplicateEmail    = errors.New("DuplicateEmailError")
	ErrValidation        = errors.New("ValidationError")
	ErrUnauthorized      = errors.New("Unauthorized")
	ErrInternalServer    = errors.New("InternalServerError")
)

func CustomizeError(err error, message string) *CustomError {
	return &CustomError{
		Err:     err,
		Message: message,
	}
}

func ValidationError(message string) *CustomError {
	return CustomizeError(ErrValidation, message)
}
