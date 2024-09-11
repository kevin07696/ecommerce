package user

import (
	"context"

	"github.com/kevin07696/ecommerce/domain/auth/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (ur UserRepository) GetUserByEmail(ctx context.Context, email models.Email) (models.User, error) {
	var user models.User
	filter := bson.D{{Key: "local", Value: email.Local}, {Key: "domain", Value: email.Domain}}
	err := ur.Collection.FindOneAndDecode(ctx, filter, &user)
	return user, err
}

func (ur UserRepository) GetUserByUsername(ctx context.Context, username models.Username) (models.User, error) {
	var user models.User
	filter := bson.D{{Key: "username", Value: username}}
	err := ur.Collection.FindOneAndDecode(ctx, filter, &user)
	return user, err
}
