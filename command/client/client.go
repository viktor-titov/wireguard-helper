package client_cmd

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "client",
		Short: "work with configuration clients",
	}

	cmd.AddCommand(newAddCommand())

	return cmd
}
