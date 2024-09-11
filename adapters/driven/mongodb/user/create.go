package user

import (
	"context"
	"log"
	"strings"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur UserRepository) CreateUser(ctx context.Context, user models.User) error {
	_, err := ur.Collection.InsertOne(ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			switch {
			case strings.Contains(err.Error(), "index: local_-1_domain_-1 dup key"):
				return domain.ErrDuplicateEmail
			case strings.Contains(err.Error(), "index: username_-1 dup key"):
				return domain.ErrDuplicateUsername
			default:
				log.Printf("%s: %v", domain.ErrDuplicateKey.Error(), err)
				return domain.ErrInternalServer
			}
		}
		return domain.ErrInternalServer
	}
	return nil
}
