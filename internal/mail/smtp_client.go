package mail

import (
	"fmt"
	"os"

	"github.com/wneessen/go-mail"
)

var sender string
var password string

func newSMTPClient() (*mail.Client, error) {
	client, err := mail.NewClient(
		"smtp.gmail.com",
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(getSender()),
		mail.WithPassword(getPassword()),
		mail.WithTLSPolicy(mail.TLSMandatory),
	)
	if err != nil {
		return nil, fmt.Errorf("create smtp client %w", err)
	}

	return client, nil
}

func getSender() string {
	if sender == "" {
		value := os.Getenv("EMAIL_SENDER")
		if value == "" {
			panic("email sender empty")
		}

		return value
	}

	return sender
}

func getPassword() string {
	if password == "" {
		v := os.Getenv("EMAIL_PASSWORD")
		if v == "" {
			panic("email password empty")
		}

		return v
	}

	return password
}
