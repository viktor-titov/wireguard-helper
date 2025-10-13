/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	client_cmd "github.com/viktor-titov/wireguard-helper/command/client"
)

var rootCmd = &cobra.Command{
	Use: "wireguard-helper",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(client_cmd.NewCommand())
}
