package email_test

import (
	"fmt"
	"testing"

	"github.com/kevin07696/ecommerce/adapters/driven/email"
	"github.com/stretchr/testify/assert"
	"github.com/wneessen/go-mail"
)

func TestNewEmailClient(t *testing.T) {
	tc := []struct {
		Name          string
		Host          string
		Opts          []mail.Option
		ExpectedError error
	}{
		{
			Name:      "Succeeds",
			Host:      "smtp.domain.tld",
			Opts:      []mail.Option{},
		},
		{
			Name:          "Fails_EmptyEmailAddr",
			Host:          "smtp.domain.tld",
			Opts:          []mail.Option{},
			ExpectedError: fmt.Errorf(email.NewEmailClientErrHeader, email.ErrMsgEmptyEmailAddr),
		},
		{
			Name:          "Fails_EmptyHost",
			Host:          "",
			Opts:          []mail.Option{},
			ExpectedError: fmt.Errorf(email.NewEmailClientErrHeader, email.ErrMsgEmptyHost),
		},
	}

	for i := range tc {
		t.Run(tc[i].Name, func(t *testing.T) {
			_, err := email.NewEmailClient(tc[i].Host, tc[i].Opts...)

			assert.Equal(t, err, tc[i].ExpectedError)
		})
	}

	tc = []struct {
		Name          string
		Host          string
		Opts          []mail.Option
		ExpectedError error
	}{
		{
			Name:          "Fails_InvalidPort",
			Host:          "smtp.domain.tld",
			Opts:          []mail.Option{mail.WithPort(65536)},
			ExpectedError: mail.ErrInvalidPort,
		},
		{
			Name:          "Fails_InvalidTimeout",
			Host:          "smtp.domain.tld",
			Opts:          []mail.Option{mail.WithTimeout(0)},
			ExpectedError: mail.ErrInvalidTimeout,
		},
	}

	for i := range tc {
		t.Run(tc[i].Name, func(t *testing.T) {
			_, err := email.NewEmailClient(tc[i].Host, tc[i].Opts...)

			assert.ErrorContains(t, err, tc[i].ExpectedError.Error())
		})
	}
}
