package services_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-session/session"
	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/mocks"
	"github.com/kevin07696/ecommerce/domain/auth/models"
	"github.com/kevin07696/ecommerce/domain/auth/port"
	"github.com/kevin07696/ecommerce/domain/auth/services"
	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	tc := []struct {
		Name          string
		Request       services.CreateSessionReq
		RefreshMock   func(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error)
		ExpectedError *domain.CustomError
	}{
		{
			Name:    "Succeeds",
			Request: services.CreateSessionReq{Username: "user"},
			RefreshMock: func(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error) {
				return &mocks.MockSession{
					SetMock: func(key string, value interface{}) {},
					SaveMock: func() error {
						return nil
					},
				}, nil
			},
			ExpectedError: nil,
		},
		{
			Name:    "FailsSaveSession_InternalServer",
			Request: services.CreateSessionReq{Username: "user"},
			RefreshMock: func(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error) {
				return &mocks.MockSession{
					SetMock: func(key string, value interface{}) {},
					SaveMock: func() error {
						return domain.ErrInternalServer
					},
				}, nil
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
		{
			Name:    "FailsRefreshSession_InternalServer",
			Request: services.CreateSessionReq{Username: "user"},
			RefreshMock: func(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error) {
				return nil, domain.ErrInternalServer
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
	}

	for i := range tc {
		t.Run(tc[i].Name, func(t *testing.T) {
			var sessionManager port.ISessionManager = &mocks.MockSessionManager{
				RefreshMock: tc[i].RefreshMock,
			}
			var cacher port.ICache
			var emailer port.IEmail
			var repositor port.IRepository
			var modeler models.IModels
			service := services.NewService(repositor, sessionManager, cacher, emailer, modeler)

			var w http.ResponseWriter
			var r *http.Request
			err := service.CreateSession(context.TODO(), w, r, tc[i].Request)

			assert.Equal(t, tc[i].ExpectedError, err)
		})
	}
}
