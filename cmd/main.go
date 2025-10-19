package main

import (
	"github.com/spf13/cobra"
	cmd "github.com/viktor-titov/wireguard-helper/internal/command"
	client_cmd "github.com/viktor-titov/wireguard-helper/internal/command/client"
)

var version string

var rootCmd = &cobra.Command{
	Use: "wireguard-helper",
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(client_cmd.NewCommand())
	rootCmd.AddCommand(cmd.NewVersionCommand(version))
}

func main() {
	Execute()
}
