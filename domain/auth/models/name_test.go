package models_test

import (
	"testing"

	"github.com/kevin07696/ecommerce/domain/auth/models"
	"github.com/stretchr/testify/assert"
)

func TestUsername(t *testing.T) {
	tc := []struct {
		Name          string
		UsernameReq   string
		UsernameResp  models.Username
		ExpectedError error
	}{
		{
			Name:          "FailsWithEmptyUsername",
			UsernameReq:   "",
			ExpectedError: models.ErrEmptyUsername,
		},
		{
			Name:          "FailsWithUsernameTooShort",
			UsernameReq:   "v1",
			ExpectedError: models.ErrInvalidUsername,
		},
		{
			Name:          "FailsWithUsernameTooLong",
			UsernameReq:   "v12345678901234567890",
			ExpectedError: models.ErrInvalidUsername,
		},
		{
			Name:          "FailsWithWithLeadingNumber",
			UsernameReq:   "12345678s",
			ExpectedError: models.ErrInvalidUsername,
		},
		{
			Name:          "FailsWithInvalidSymbol",
			UsernameReq:   "user$1",
			ExpectedError: models.ErrInvalidUsername,
		},
		{
			Name:          "Succeeds",
			UsernameReq:   "userv1",
			UsernameResp:  "userv1",
			ExpectedError: nil,
		},
		{
			Name:          "SucceedsWithCaps",
			UsernameReq:   "UserV1",
			UsernameResp:  "UserV1",
			ExpectedError: nil,
		},
		{
			Name:          "SucceedsWith_",
			UsernameReq:   "user_v1",
			UsernameResp:  "user_v1",
			ExpectedError: nil,
		},
	}

	for i := range tc {
		t.Run(tc[i].Name, func(t *testing.T) {
			models := models.Models{}
			un, err := models.NewUsername(tc[i].UsernameReq)

			assert.Equal(t, tc[i].ExpectedError, err)
			assert.Equal(t, tc[i].UsernameResp, un)
		})
	}
}
