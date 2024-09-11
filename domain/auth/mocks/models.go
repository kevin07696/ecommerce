package mocks

import (
	"github.com/kevin07696/ecommerce/domain/auth/models"
)

type MockModels struct {
	NewEmailMock    func(email string) (models.Email, error)
	NewUsernameMock func(username string) (models.Username, error)
	NewOTPMock      func(otp string) (models.OTP, error)
	NewRoleMock     func(role string) (models.Role, error)
	NewUserMock     func(username models.Username, email models.Email, role models.Role) models.User
}

func (m MockModels) NewEmail(email string) (models.Email, error) {
	return m.NewEmailMock(email)
}

func (m MockModels) NewUsername(username string) (models.Username, error) {
	return m.NewUsernameMock(username)
}

func (m MockModels) NewOTP(otp string) (models.OTP, error) {
	return m.NewOTPMock(otp)
}

func (m MockModels) NewRole(role string) (models.Role, error) {
	return m.NewRoleMock(role)
}

func (m MockModels) NewUser(username models.Username, email models.Email, role models.Role) models.User {
	return m.NewUserMock(username, email, role)
}
