package services

const (
	ErrMsgUsernameNotUnique = "Sorry, that username is taken. Please try again."
	ErrMsgEmailNotUnique    = "That email is already registered. Please log in or use a different email."
	ErrMsgEmailNotFound     = "Sorry, we couldn't find an account with that email. Please check your spelling and try again"
	ErrMsgUserIdNotFound    = "Sorry, we couldn't find an account with that username or email. Please check your spelling and try again"
	ErrMsgUnauthorized      = "Session has expired. Please log in."
	ErrMsgInternalServer    = "Sorry, server error. We're fixing it."

	ErrMsgEmptyEmail      = "Email cannot be empty."
	ErrMsgInvalidEmail    = "Email is invalid. Please verify with the provided hint and try again."
	ErrMsgEmptyUsername   = "Username cannot be empty."
	ErrMsgInvalidUsername = "Username is invalid. Please verify with the provided hint and try again."
	ErrMsgEmptyUserId     = "Email or username cannot be empty."
	ErrMsgInvalidUserId   = "Please provide a valid username or email."
	ErrMsgEmptyOTP        = "One Time Password cannot be empty"
	ErrMsgInvalidOTP      = "One Time Password is incorrect"
	ErrMsgEmptyRole       = "Role cannot be empty"
	ErrMsgInvalidRole     = "Role is invalid"
)
