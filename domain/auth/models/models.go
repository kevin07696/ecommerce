package models

type IModels interface {
	NewEmail(email string) (Email, error)
	NewUsername(username string) (Username, error)
	NewOTP(otp string) (OTP, error)
	NewRole(role string) (Role, error)
	NewUser(username Username, email Email, role Role) User
}

type Models struct{}
