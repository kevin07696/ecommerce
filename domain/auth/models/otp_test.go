package models_test

import (
	"testing"

	"github.com/kevin07696/ecommerce/domain/auth/models"
	"github.com/stretchr/testify/assert"
)

func TestNewOTP(t *testing.T) {
	tc := []struct {
		Name          string
		OTPReq        string
		OTPResp       models.OTP
		ExpectedError error
	}{
		{
			Name:    "Succeeds",
			OTPReq:  "S@1$fz_i",
			OTPResp: "S@1$fz_i",
		},
		{
			Name:          "FailsWithEmptyOTP",
			OTPReq:        "",
			ExpectedError: models.ErrEmptyOTP,
		},
		{
			Name:          "FailsWithInvalidOTP",
			OTPReq:        "123456789",
			ExpectedError: models.ErrInvalidOTP,
		},
	}

	for i := range tc {
		t.Run(tc[i].Name, func(t *testing.T) {
			models := models.Models{}
			otp, err := models.NewOTP(tc[i].OTPReq)

			assert.Equal(t, tc[i].ExpectedError, err)
			assert.Equal(t, tc[i].OTPResp, otp)
		})
	}
}
