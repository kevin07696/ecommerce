package services_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/mocks"
	"github.com/kevin07696/ecommerce/domain/auth/models"
	"github.com/kevin07696/ecommerce/domain/auth/port"
	"github.com/kevin07696/ecommerce/domain/auth/services"
	"github.com/stretchr/testify/assert"
)

func TestDeleteSession(t *testing.T) {
	tc := []struct {
		Name          string
		DeleteMock    func(ctx context.Context, w http.ResponseWriter, r *http.Request) error
		ExpectedError *domain.CustomError
	}{
		{
			Name: "Succeeds",
			DeleteMock: func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
				return nil
			},
			ExpectedError: nil,
		},
		{
			Name: "FailsDeleteSession_InternalServer",
			DeleteMock: func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
				return domain.ErrInternalServer
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
	}

	for i := range tc {
		t.Run(tc[i].Name, func(t *testing.T) {
			var sessionManager port.ISessionManager = &mocks.MockSessionManager{
				DeleteMock: tc[i].DeleteMock,
			}
			var cacher port.ICache
			var emailer port.IEmail
			var repositor port.IRepository
			var modeler models.IModels
			service := services.NewService(repositor, sessionManager, cacher, emailer, modeler)

			var w http.ResponseWriter
			var r *http.Request
			err := service.DeleteSession(context.TODO(), w, r)

			assert.Equal(t, tc[i].ExpectedError, err)
		})
	}
}
