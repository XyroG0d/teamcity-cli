/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/XyroG0d/teamcity-cli/configManager"
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
		//TODO: need to perform below logic when --list argument is passed
		listFlag, _ := cmd.Flags().GetBool("list")
		if len(args) == 0 {
			if listFlag {
				configSettings := configManager.ReadConfigMap()
				for k, v := range configSettings {
					fmt.Printf("%s=%s\n", k, v)
				}
			} else {
				fmt.Println(cmd.Help())
			}
		} else {
			configSettings := configManager.ReadConfigMap()
			if len(args) == 1 {
				dat := args[0] + "=" + configSettings[args[0]]
				fmt.Println(dat) //TODO: need to add error handling if flag is not valid
			} else if len(args) == 2 {
				configSettings[args[0]] = args[1]
				configManager.WriteConfig(configSettings)
			} else {
				panic("Too many arguments") // replace with print and exit command
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configManager.CreateConfigFileIfNotExists()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	configCmd.Flags().BoolP("list", "l", false, "Displays all configurations")
}
