package configManager

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var teamcityURL string = ""

var gitBranch string = ""

var Interval int = -100

var Until int = -100

var authToken string = ""

const configFile = "tctcli.txt"

const DEFAULT_GIT_BRANCH = "main"

const DEFAULT_INTERVAL = 5000
const DEFAULT_UNTIL = 50000

func IntializeConfig() {
	if gitBranch == "" {
		gitBranch = DEFAULT_GIT_BRANCH
	}
	if Interval == -100 { //is there a better way to do this?
		Interval = DEFAULT_INTERVAL
	}
	if Until == -100 {
		Interval = DEFAULT_UNTIL
	}

	configMap := ReadConfigMap()
	teamcityURL = configMap["teamCityURL"]
	gitBranch = configMap["gitBranch"]
	authToken = configMap["authToken"]
	if configMap["interval"] != "" {
		res, err := strconv.Atoi(configMap["interval"])
		if err == nil {
			Interval = res
		} else {
			fmt.Println("blabla", configMap["interval"])
			panic(err)
		}
	}
	if configMap["until"] != "" {
		res, err := strconv.Atoi(configMap["until"])
		if err == nil {
			Until = res
		} else {
			panic(err)
		}
	}
	authToken = "abc" //TODO: temp assignment, need to read from file instead
	if authToken == "" {
		panic("Auth token has to be set, please use tct-cli config authToken <authToken> to set it")
	}
}

// initalizing static map for now, but this needs to be read from a file to maintain persistance
func ReadConfigMap() map[string]string {
	var file *os.File
	var err error
	configMap := make(map[string]string)

	file, err = os.Open(configFile)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines) //not getting correct results here probably due to different character being replaced at EOL
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "=")
		value := ""
		key := strings.Trim(split[0], " ")
		if len(split) == 2 {
			value = split[1]
		}
		configMap[key] = value
	}

	return configMap
}

func WriteConfig(configMap map[string]string) {
	file, err := os.OpenFile(configFile, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		panic(err)
	}
	for k, v := range configMap {
		dat := k + "=" + v + "\n"
		_, err := file.WriteString(dat)
		if err != nil {
			panic(err)
		}
	}
}

func CreateConfigFileIfNotExists() {
	_, err := os.Open(configFile)
	if err != nil {
		os.Create(configFile)
		configMap := make(map[string]string)
		configMap["teamCityURL"] = ""
		configMap["gitBranch"] = DEFAULT_GIT_BRANCH
		configMap["interval"] = "5000"
		configMap["until"] = "20000"
		WriteConfig(configMap)
	}
}
