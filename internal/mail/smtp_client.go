package mail

import (
	"fmt"

	"github.com/wneessen/go-mail"
)

const (
	sender = "kkeeptorch@gmail.com"
	pass   = "dzjx jzlr hdlf mijo"
)

func newSMTPClient() (*mail.Client, error) {
	client, err := mail.NewClient(
		"smtp.gmail.com",
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(sender),
		mail.WithPassword(pass),
		mail.WithTLSPolicy(mail.TLSMandatory),
	)
	if err != nil {
		return nil, fmt.Errorf("create smtp client %w", err)
	}

	return client, nil
}
