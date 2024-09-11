package email

import (
	"context"
	"log/slog"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/wneessen/go-mail"
)

func (e Wrapper) SendEmail(ctx context.Context, to, subject, body string) error {
	m := mail.NewMsg()
	if err := m.From(e.EmailAddr); err != nil {
		slog.LogAttrs(ctx, slog.LevelError, "em.SendEmail: failed to set From address", slog.String("From", e.EmailAddr), slog.Any("error", err))
		return domain.ErrInternalServer
	}
	if err := m.To(to); err != nil {
		slog.LogAttrs(ctx, slog.LevelError, "em.SendEmail: failed to set To address", slog.String("To", to), slog.Any("error", err))
		return domain.ErrInternalServer
	}
	m.Subject(subject)
	m.SetBodyString(mail.TypeTextPlain, body)

	if err := e.Client.DialAndSendWithContext(ctx, m); err != nil {
		slog.LogAttrs(ctx, slog.LevelError, "em.SendEmail: failed to send email", slog.Any("error", err))
		return domain.ErrInternalServer
	}
	slog.LogAttrs(ctx, slog.LevelDebug, "em.SendEmail: succeeded to send email")
	return nil
}
