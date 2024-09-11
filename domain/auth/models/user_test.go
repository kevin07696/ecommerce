package models_test

import (
	"testing"

	"github.com/kevin07696/ecommerce/domain/auth/models"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	tc := []struct {
		Name             string
		Username         models.Username
		Email            models.Email
		Role             models.Role
		ExpectedResponse models.User
	}{
		{
			Name:     "Succeeds",
			Username: "username",
			Email: models.Email{
				Local:      "local",
				SubAddress: "+sub",
				Domain:     "domain.sub.tld",
			},
			Role: "developer",
			ExpectedResponse: models.User{
				Username: "username",
				Email: models.Email{
					Local:      "local",
					SubAddress: "+sub",
					Domain:     "domain.sub.tld",
				},
				Role: "developer",
			},
		},
	}

	for i := range tc {
		t.Run(tc[i].Name, func(t *testing.T) {
			models := models.Models{}
			user := models.NewUser(tc[i].Username, tc[i].Email, tc[i].Role)

			assert.Equal(t, tc[i].ExpectedResponse, user)
		})
	}
}
