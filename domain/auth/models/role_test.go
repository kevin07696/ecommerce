package models_test

import (
	"testing"

	"github.com/kevin07696/ecommerce/domain/auth/models"
	"github.com/stretchr/testify/assert"
)

func TestRole(t *testing.T) {
	tc := []struct {
		Name          string
		RoleReq       string
		RoleResp      models.Role
		ExpectedError error
	}{
		{
			Name:     "SucceedsWithVendor",
			RoleReq:  "vendor",
			RoleResp: "vendor",
		},
		{
			Name:     "SucceedsWithShopper",
			RoleReq:  "shopper",
			RoleResp: "shopper",
		},
		{
			Name:     "SucceedsWithDeveloper",
			RoleReq:  "developer",
			RoleResp: "developer",
		},
		{
			Name:          "FailsWithEmptyRole",
			RoleReq:       "",
			ExpectedError: models.ErrEmptyRole,
		},
		{
			Name:          "FailsWithInvalidRole",
			RoleReq:       "Admin",
			ExpectedError: models.ErrInvalidRole,
		},
		{
			Name:          "FailsWithCaseSensitiveRole",
			RoleReq:       "Developer",
			ExpectedError: models.ErrInvalidRole,
		},
	}

	for i := range tc {
		t.Run(tc[i].Name, func(t *testing.T) {
			models := models.Models{}
			role, err := models.NewRole(tc[i].RoleReq)

			assert.Equal(t, tc[i].ExpectedError, err)
			assert.Equal(t, tc[i].RoleResp, role)
		})
	}
}
