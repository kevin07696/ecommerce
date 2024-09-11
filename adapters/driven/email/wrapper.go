package email

import "fmt"

type Wrapper struct {
	Client    IClient
	EmailAddr string
}

func NewClientWrapper(client IClient, emailAddr string) (Wrapper, error) {
	if emailAddr == "" {
		return Wrapper{}, fmt.Errorf(NewEmailClientErrHeader, ErrMsgEmptyEmailAddr)
	}
	return Wrapper{
		Client:    client,
		EmailAddr: emailAddr,
	}, nil
}
