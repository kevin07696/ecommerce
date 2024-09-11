package mocks

import (
	"context"

	"github.com/kevin07696/ecommerce/domain/auth/models"
)

type MockRepository struct {
	GetUserByEmailMock    func(ctx context.Context, email models.Email) (models.User, error)
	GetUserByUsernameMock func(ctx context.Context, username models.Username) (models.User, error)
	CreateUserMock        func(ctx context.Context, user models.User) error
}

func (m MockRepository) GetUserByEmail(ctx context.Context, email models.Email) (models.User, error) {
	return m.GetUserByEmailMock(ctx, email)
}

func (m MockRepository) GetUserByUsername(ctx context.Context, username models.Username) (models.User, error) {
	return m.GetUserByUsernameMock(ctx, username)
}

func (m MockRepository) CreateUser(ctx context.Context, user models.User) error {
	return m.CreateUserMock(ctx, user)
}
