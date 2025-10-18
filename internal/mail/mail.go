package mail

import (
	"fmt"

	"github.com/wneessen/go-mail"
)

func Send(recipient, subject string, attach []string) (string, error) {
	message := mail.NewMsg()

	err := message.From(sender)
	if err != nil {
		return "", fmt.Errorf("setting up sender %w", err)
	}

	err = message.To(recipient)
	if err != nil {
		return "", fmt.Errorf("setting up recipient %w", err)
	}

	message.Subject(subject)
	body, err := templateMail("testmail")
	if err != nil {
		return "", fmt.Errorf("make body %w", err)
	}

	message.SetBodyString(mail.TypeTextPlain, body.String())

	// Добавление вложения
	for _, path := range attach {
		message.AttachFile(path)
	}

	client, err := newSMTPClient()
	if err != nil {
		return "", fmt.Errorf("get new smtp client %w", err)
	}

	err = client.DialAndSend(message)
	if err != nil {
		return "", fmt.Errorf("send mail %w", err)
	}

	return fmt.Sprintf("The email was successfully sent to %s", recipient), nil
}
