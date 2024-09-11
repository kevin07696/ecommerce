package models_test

import (
	"testing"

	"github.com/kevin07696/ecommerce/domain/auth/models"
	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	tc := []struct {
		Name          string
		EmailReq      string
		EmailResp     models.Email
		ExpectedError error
	}{
		{
			Name:     "Succeeds",
			EmailReq: "lo.cal@domain.tld",
			EmailResp: models.Email{
				Local:      "lo.cal",
				SubAddress: "",
				Domain:     "domain.tld",
			},
		},
		{
			Name:     "Succeeds_+SubAddr",
			EmailReq: "lo.cal+sub@domain.tld",
			EmailResp: models.Email{
				Local:      "lo.cal",
				SubAddress: "+sub",
				Domain:     "domain.tld",
			},
		},
		{
			Name:     "Succeeds_-SubAddr",
			EmailReq: "lo.cal-sub@domain.tld",
			EmailResp: models.Email{
				Local:      "lo.cal",
				SubAddress: "-sub",
				Domain:     "domain.tld",
			},
		},
		{
			Name:     "Succeeds_SubDomain",
			EmailReq: "lo.cal@domain.sub.tld",
			EmailResp: models.Email{
				Local:      "lo.cal",
				SubAddress: "",
				Domain:     "domain.sub.tld",
			},
		},
		{
			Name:          "Fails_EmptyEmail",
			EmailReq:      "",
			ExpectedError: models.ErrEmptyEmail,
		},
		{
			Name:          "Fails_InvalidLocalSymbol",
			EmailReq:      "local+$ub@domain.com",
			ExpectedError: models.ErrInvalidEmail,
		},
		{
			Name:          "Fails_Domain",
			EmailReq:      "lo.cal@domain",
			ExpectedError: models.ErrInvalidEmail,
		},
	}

	for i := range tc {
		t.Run("NewEmail/"+tc[i].Name, func(t *testing.T) {
			models := models.Models{}
			em, err := models.NewEmail(tc[i].EmailReq)

			assert.Equal(t, tc[i].ExpectedError, err)
			assert.Equal(t, tc[i].EmailResp, em)
		})
		t.Run("ToString/"+tc[i].Name, func(t *testing.T) {
			if tc[i].ExpectedError == nil {
				email := tc[i].EmailResp.ToString()

				assert.Equal(t, tc[i].EmailReq, email)
			}
		})
	}
}
