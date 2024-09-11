package email

import (
	"context"
	"fmt"
	"log"

	"github.com/wneessen/go-mail"
)

type IClient interface {
	DialAndSendWithContext(ctx context.Context, messages ...*mail.Msg) error
}

const NewEmailClientErrHeader = "NewEmailClientError: %v"

func NewEmailClient(host string, opts ...mail.Option) (IClient, error) {
	if host == "" {
		return nil, fmt.Errorf(NewEmailClientErrHeader, ErrMsgEmptyHost)
	}
	emailer, err := mail.NewClient(host, opts...)
	if err != nil {
		return nil, fmt.Errorf(NewEmailClientErrHeader, err)
	}
	log.Printf("Connecting to GoMail client: %s", emailer.ServerAddr())
	return emailer, nil
}
