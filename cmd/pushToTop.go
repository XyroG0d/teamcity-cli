/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/XyroG0d/teamcity-cli/configManager"
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
		interval, _ := cmd.Flags().GetInt("interval")

		until, _ := cmd.Flags().GetInt("until")

		if interval == 0 && until != 0 {
			interval = configManager.DEFAULT_INTERVAL
		}

		if until == 0 && interval != 0 {
			until = configManager.DEFAULT_UNTIL
		}

		if interval > until {
			fmt.Println("Interval cannot be more than Until")
			os.Exit(1)
		}
		buildNo, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Argument passed is not a number")
			os.Exit(1)
		}
		pushBuildToTopAction := pushBuildToTop

		if interval == 0 && until == 0 {
			buildNo, err := strconv.Atoi(args[0])
			if err != nil {
				pushBuildToTopAction(buildNo)
			}
		} else {
			repeatPushBuildToTop(pushBuildToTopAction, buildNo, interval, until)
		}
	},
}

func repeatPushBuildToTop(pushBuildAction func(int), buildNumber int, interval int, until int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Millisecond)
	startTime := time.Now()
	go func() {
		for range ticker.C {
			if time.Since(startTime) >= time.Duration(until)*time.Millisecond {
				os.Exit(1)
			}
			pushBuildAction(buildNumber)
		}
	}()

	time.Sleep(time.Duration(until) * time.Second)
	ticker.Stop()
}

func pushBuildToTop(buildNumber int) {
	fmt.Println("Hi")
	//TODO: call API
}

func init() {
	rootCmd.AddCommand(pushToTopCmd)
	configManager.IntializeConfig()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushToTopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	pushToTopCmd.Flags().Int("interval", 0, "Interval Flag") //need to add more context
	pushToTopCmd.Flags().Int("until", 0, "Until Flag")       //need to add more context
}
