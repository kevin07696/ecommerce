package port

import "context"

type IEmail interface {
	SendEmail(ctx context.Context, to, subject, body string) error
}
