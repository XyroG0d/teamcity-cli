package config

import (
	"strconv"
)

var teamcityURL string = ""

var gitBranch string = ""

var interval int = -100

var until int = -100

var authToken string = ""

const DEFAULT_GIT_BRANCH = "main"

const DEFAULT_INTERVAL = 5000
const DEFAULT_UNTIL = 50000

func IntializeConfig() {
	if gitBranch == "" {
		gitBranch = DEFAULT_GIT_BRANCH
	}
	if interval == -100 { //is there a better way to do this?
		interval = DEFAULT_INTERVAL
	}
	if until == -100 {
		interval = DEFAULT_UNTIL
	}

	configMap := loadConfig()
	teamcityURL = configMap["teamCityURL"]
	gitBranch = configMap["gitBranch"]
	authToken = configMap["authToken"]
	res, err := strconv.Atoi(configMap["interval"])
	if err == nil {
		interval = res
	} else {
		panic(err)
	}
	res1, err1 := strconv.Atoi(configMap["until"])
	if err1 == nil {
		until = res1
	} else {
		panic(err1)
	}
	if authToken == "" {
		panic("Auth token has to be set, please use tct-cli config authToken <authToken> to set it")
	}
}

func loadConfig() map[string]string {
	//initalizing static map for now, but this needs to be read from a file to maintain persistance
	configMap := make(map[string]string)
	configMap["teamCityURL"] = "someURL"
	configMap["gitBranch"] = "Develop"
	configMap["gitBranch"] = "Develop"
	configMap["interval"] = ""
	configMap["until"] = ""
	return configMap
}
