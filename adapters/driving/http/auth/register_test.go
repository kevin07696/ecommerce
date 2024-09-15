package auth_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	hauth "github.com/kevin07696/ecommerce/adapters/driving/http/auth"
	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/mocks"
	sauth "github.com/kevin07696/ecommerce/domain/auth/services"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {

	tc := []struct {
		Name              string
		RegisterUserMock  func(ctx context.Context, req sauth.RegisterUserReq) (sauth.RegisterUserResp, *domain.CustomError)
		CreateSessionMock func(ctx context.Context, w http.ResponseWriter, r *http.Request, req sauth.CreateSessionReq) *domain.CustomError
		ExpectedBody      string
		// ExpectedStatus    int
	}{
		{
			Name: "Fails_EmptyUsername",
			RegisterUserMock: func(ctx context.Context, req sauth.RegisterUserReq) (sauth.RegisterUserResp, *domain.CustomError) {
				return sauth.RegisterUserResp{}, domain.ValidationError(sauth.ErrMsgEmptyUsername)
			},
			ExpectedBody: `<div>Username cannot be empty.</div>`,
			// ExpectedStatus: 400,
		},
		{
			Name: "Fails_InvalidUsername",
			RegisterUserMock: func(ctx context.Context, req sauth.RegisterUserReq) (sauth.RegisterUserResp, *domain.CustomError) {
				return sauth.RegisterUserResp{}, domain.ValidationError(sauth.ErrMsgInvalidUsername)
			},
			ExpectedBody: `<div>Username is invalid. Please verify with the provided hint and try again.</div>`,
			// ExpectedStatus: 400,
		},
		{
			Name: "Fails_EmptyEmail",
			RegisterUserMock: func(ctx context.Context, req sauth.RegisterUserReq) (sauth.RegisterUserResp, *domain.CustomError) {
				return sauth.RegisterUserResp{}, domain.ValidationError(sauth.ErrMsgEmptyEmail)
			},
			ExpectedBody: `<div>Email cannot be empty.</div>`,
			// ExpectedStatus: 400,
		},
		{
			Name: "Fails_InvalidEmail",
			RegisterUserMock: func(ctx context.Context, req sauth.RegisterUserReq) (sauth.RegisterUserResp, *domain.CustomError) {
				return sauth.RegisterUserResp{}, domain.ValidationError(sauth.ErrMsgInvalidEmail)
			},
			ExpectedBody: `<div>Email is invalid. Please verify with the provided hint and try again.</div>`,
			// ExpectedStatus: 400,
		},
	}

	for i := range tc {
		t.Run(tc[i].Name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/api/create-account", nil)
			if err != nil {
				t.Fatal(err)
			}

			api := &mocks.MockAPI{}
			api.RegisterUserMock = tc[i].RegisterUserMock

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(hauth.HandleCreateAccount(api))
			handler.ServeHTTP(rr, req)

			assert.Equal(t, tc[i].ExpectedBody, rr.Body.String())
			// assert.Equal(t, tc[i].ExpectedStatus, rr.Code)
		})
	}

}
