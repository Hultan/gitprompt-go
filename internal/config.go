package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path"
)

type Config struct {
	Format string `json:format`
	IncludeZeroValues bool `json:includeZeroValues`
	Separator string `json:separator`
	PromptPrefix string `json:promptPrefix`
	PromptSuffix string `json:promptSuffix`
	Branch struct {
		Prefix string `json:prefix`
		Suffix string `json:suffix`
	} `json:branch`
	Ahead struct {
		Prefix string `json:prefix`
		Suffix string `json:suffix`
	} `json:ahead`
	Behind struct {
		Prefix string `json:prefix`
		Suffix string `json:suffix`
	} `json:behind`
	Untracked struct {
		Prefix string `json:prefix`
		Suffix string `json:suffix`
	} `json:untracked`
	Modified struct {
		Prefix string `json:prefix`
		Suffix string `json:suffix`
	} `json:modified`
	Deleted struct {
		Prefix string `json:prefix`
		Suffix string `json:suffix`
	} `json:deleted`
	Unmerged struct {
		Prefix string `json:prefix`
		Suffix string `json:suffix`
	} `json:unmerged`
	Staged struct {
		Prefix string `json:prefix`
		Suffix string `json:suffix`
	} `json:staged`
}

func NewConfig() *Config {
	return new(Config)
}

func (config *Config) ConfigExists() bool {
	// Get the path to the config file
	path := getConfigPath()

	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Load : Loads the configuration file
func (config *Config) Load() error {
	// Get the path to the config file
	path := getConfigPath()

	// Open config file
	configFile, err := os.Open(path)

	// Handle errors
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()

	// Parse the JSON document
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return nil
}


// Get path to the config file
func getConfigPath() string {
	home := getHomeDirectory()

	return path.Join(home, ".config/gitprompt-go/config.json")
}

// Get current users home directory
func getHomeDirectory() string {
	u, err := user.Current()
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get user home directory : %s", err)
		panic(errorMessage)
	}
	return u.HomeDir
}
