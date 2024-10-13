/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pushToTopCmd represents the pushToTop command
var pushToTopCmd = &cobra.Command{
	Use:   "pushToTop",
	Short: "This command will move the build with given build number to the top of its queue",
	Long: `This command will move the build with given build number to the top of its queue

Usage: tct-cli moveToTop <buildNumber> [--interval] <timeInterval> [--until] <untilTime>

It also take two optional flags --interval and --until.

If neither are provided then this command will only move the build to top once.

If either are provided, the moveToTop command will run internally till <unTilTime> milliseconds at an interval of <timeInternal> milliseconds

Default Values:
<timeInterval> : 5000
<takeUntil> : 50000
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Need to call pushToTop API here")
	},
}

func init() {
	rootCmd.AddCommand(pushToTopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushToTopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushToTopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
