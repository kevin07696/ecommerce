package port

import (
	"context"

	"github.com/kevin07696/ecommerce/domain/auth/models"
)

type IRepository interface {
	GetUserByEmail(ctx context.Context, email models.Email) (models.User, error)
	GetUserByUsername(ctx context.Context, username models.Username) (models.User, error)
	CreateUser(context.Context, models.User) error
}
