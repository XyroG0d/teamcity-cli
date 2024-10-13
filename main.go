/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/XyroG0d/teamcity-cli/cmd"
	"github.com/XyroG0d/teamcity-cli/config"
)

func main() {
	cmd.Execute()
	config.IntializeConfig()
}
