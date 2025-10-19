package cmd

import "github.com/spf13/cobra"

func NewVersionCommand(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Version of wireguard helper",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("%s\n", version)
		},
	}

	return cmd
}
