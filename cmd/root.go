/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	client_cmd "github.com/viktor-titov/wireguard-helper/internal/command/client"
)

var rootCmd = &cobra.Command{
	Use: "wireguard-helper",
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(client_cmd.NewCommand())
}
