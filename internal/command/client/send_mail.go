package client_cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/viktor-titov/wireguard-helper/internal/mail"
)

func newSendMailCommand() *cobra.Command {
	var recipientMail string
	var configs []string

	cmd := &cobra.Command{
		Use:   "send",
		Short: "Send mail with config to client",
		RunE: func(cmd *cobra.Command, args []string) error {
			if recipientMail == "" {
				return errors.New("Mail of recipient must be provided via --recipient")
			}

			if configs == nil {
				return errors.New("Specify configuration files to transfer via --config or -c")
			}

			msg, err := mail.Send(
				recipientMail,
				"Конфигурация клиента для wireguard",
				configs,
			)
			if err != nil {
				return err
			}

			cmd.Println(msg)

			return nil
		},
	}

	cmd.Flags().StringVarP(&recipientMail, "recipient", "r", "", "Mail of recipient of config wireguard")
	cmd.Flags().StringArrayVarP(&configs, "config", "c", nil, "Config for sending client")

	return cmd
}
