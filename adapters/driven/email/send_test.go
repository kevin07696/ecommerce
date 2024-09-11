package email_test

import (
	"context"
	"errors"
	"testing"

	"github.com/kevin07696/ecommerce/adapters/driven/email"
	"github.com/kevin07696/ecommerce/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wneessen/go-mail"
)

type mockClient struct {
	mock.Mock
}

func (m *mockClient) DialAndSendWithContext(ctx context.Context, messages ...*mail.Msg) error {
	args := m.Called(ctx, messages)
	return args.Error(0)
}

func TestSendEmail(t *testing.T) {
	tc := []struct {
		Name             string
		From             string
		To               string
		Subject          string
		Body             string
		DialAndSendError error
		ExpectedError    error
	}{
		{
			Name:    "Succeeds",
			From:    "local+sub@domain.sub.tld",
			To:      "local+sub@domain.sub.tld",
			Subject: "subject",
			Body:    "body",
		},
		{
			Name:          "Fails_InvalidFrom",
			From:          "local+subdomain.sub.tld",
			To:            "local+sub@domain.sub.tld",
			Subject:       "subject",
			Body:          "body",
			ExpectedError: domain.ErrInternalServer,
		},
		{
			Name:          "Fails_InvalidTo",
			From:          "local+sub@domain.sub.tld",
			To:            "local+subdomain.sub.tld",
			Subject:       "subject",
			Body:          "body",
			ExpectedError: domain.ErrInternalServer,
		},
		{
			Name:             "Fails_DialAndSend",
			From:             "local+sub@domain.sub.tld",
			To:               "local+sub@domain.sub.tld",
			Subject:          "subject",
			Body:             "body",
			DialAndSendError: errors.New("Dial Failed"),
			ExpectedError:    domain.ErrInternalServer,
		},
	}

	for i := range tc {
		t.Run(tc[i].Name, func(t *testing.T) {
			mockClient := new(mockClient)
			mockClient.On("DialAndSendWithContext", mock.Anything, mock.Anything).Return(tc[i].DialAndSendError)

			emailer := email.Wrapper{
				Client:    mockClient,
				EmailAddr: tc[i].From,
			}

			err := emailer.SendEmail(context.TODO(), tc[i].To, tc[i].Subject, tc[i].Body)

			assert.Equal(t, tc[i].ExpectedError, err)
		})
	}
}
