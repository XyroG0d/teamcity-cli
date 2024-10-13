/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Check and change configuration settings",
	Long: ` Check and change configuration settings

Usage: tct-cli config [--list] [configName] <configValue>

When configName is provided, this command returns the specific value of the config specified.

If configName and configValue are provided, this command sets the value to the given config.

If --list is provided, this command lists down all the configurations
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Need to intialize a list of configs to set and show then to user")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
