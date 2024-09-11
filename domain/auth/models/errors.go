package models

import "errors"

var (
	ErrEmptyEmail      = errors.New("ErrEmptyEmail")
	ErrInvalidEmail    = errors.New("ErrInvalidEmail")
	ErrEmptyUsername   = errors.New("ErrEmptyUsername")
	ErrInvalidUsername = errors.New("ErrInvalidUsername")
	ErrEmptyOTP        = errors.New("ErrEmptyOTP")
	ErrInvalidOTP      = errors.New("ErrInvalidOTP")
	ErrEmptyRole       = errors.New("ErrEmptyRole")
	ErrInvalidRole     = errors.New("ErrInvalidRole")
)
