package mocks

import "context"

type MockEmailer struct {
	SendEmailMock func(ctx context.Context, to, subject, body string) error
}

func (e MockEmailer) SendEmail(ctx context.Context, to, subject, body string) error {
	return e.SendEmailMock(ctx, to, subject, body)
}
