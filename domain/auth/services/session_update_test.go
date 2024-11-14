package services_test

import (
	"context"
	"errors"
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

func TestUpdateSession(t *testing.T) {
	tc := []struct {
		Name          string
		sessionMock   mocks.MockSessionManager
		ExpectedError *domain.CustomError
	}{
		{
			Name: "Succeeds",
			sessionMock: mocks.MockSessionManager{
				StartMock: func(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error) {
					return &mocks.MockSession{
						GetMock: func(key string) (interface{}, bool) {
							return nil, true
						},
						SaveMock: func() error {
							return nil
						},
					}, nil
				},
				NeedsRefreshMock: func(session session.Store) bool {
					return false
				},
			},
			ExpectedError: nil,
		},
		{
			Name: "FailsStartSession_InternalServer",
			sessionMock: mocks.MockSessionManager{
				StartMock: func(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error) {
					return nil, domain.ErrInternalServer
				},
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
		{
			Name: "FailsGetSession_Unauthorized",
			sessionMock: mocks.MockSessionManager{
				StartMock: func(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error) {
					return &mocks.MockSession{
						GetMock: func(key string) (interface{}, bool) {
							return nil, false
						},
					}, nil
				},
				DeleteMock: func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
					return nil
				},
			},
			ExpectedError: domain.CustomizeError(domain.ErrUnauthorized, services.ErrMsgUnauthorized),
		},
		{
			Name: "FailsDeleteSession_InternalServer",
			sessionMock: mocks.MockSessionManager{
				StartMock: func(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error) {
					return &mocks.MockSession{
						GetMock: func(key string) (interface{}, bool) {
							return nil, false
						},
					}, nil
				},
				DeleteMock: func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
					return domain.ErrInternalServer
				},
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
		{
			Name: "RefreshSession_Succeeds",
			sessionMock: mocks.MockSessionManager{
				StartMock: func(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error) {
					return &mocks.MockSession{
						GetMock: func(key string) (interface{}, bool) {
							return nil, true
						},
					}, nil
				},
				NeedsRefreshMock: func(session session.Store) bool {
					return true
				},
				RefreshMock: func(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error) {
					return &mocks.MockSession{
						SetMock: func(key string, value interface{}) {},
						SaveMock: func() error {
							return nil
						},
					}, nil
				},
			},
			ExpectedError: nil,
		},
		{
			Name: "FailsRefreshSession_InternalServer",
			sessionMock: mocks.MockSessionManager{
				StartMock: func(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error) {
					return &mocks.MockSession{
						GetMock: func(key string) (interface{}, bool) {
							return &mocks.MockSession{}, true
						},
					}, nil
				},
				NeedsRefreshMock: func(session session.Store) bool {
					return true
				},
				RefreshMock: func(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error) {
					return nil, domain.ErrInternalServer
				},
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
		{
			Name: "FailsSaveSession_InternalServer",
			sessionMock: mocks.MockSessionManager{
				StartMock: func(ctx context.Context, w http.ResponseWriter, r *http.Request) (session.Store, error) {
					return &mocks.MockSession{
						GetMock: func(key string) (interface{}, bool) {
							return nil, true
						},
						SaveMock: func() error {
							return errors.New("")
						},
					}, nil
				},
				NeedsRefreshMock: func(session session.Store) bool {
					return false
				},
			},
			ExpectedError: domain.CustomizeError(domain.ErrInternalServer, services.ErrMsgInternalServer),
		},
	}

	for i := range tc {
		t.Run(tc[i].Name, func(t *testing.T) {
			var sessionManager port.ISessionManager = &tc[i].sessionMock
			var cacher port.ICache
			var emailer port.IEmail
			var repositor port.IRepository
			var modeler models.IModels
			service := services.NewService(repositor, sessionManager, cacher, emailer, modeler)

			var w http.ResponseWriter
			var r *http.Request
			err := service.UpdateSession(context.TODO(), w, r)

			assert.Equal(t, tc[i].ExpectedError, err)
		})
	}
}
