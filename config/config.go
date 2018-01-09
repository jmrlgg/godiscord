package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Token Public variables
var (
	Token        string // GODiscord App Token
	BotPrefix    string // BotPrefix GODiscord Bot Prefix
	ClientID     string // Reddit APP Client ID
	ClientSecret string // Reddit APP Client Secret
	RedditUser   string // Reddit UserName
	RedditPass   string // Reddit Password

	// Private variables
	config *configStruct
)

type configStruct struct {
	Token        string `json:"Token"`
	BotPrefix    string `json:"BotPrefix"`
	ClientID     string `json:"ClientID"`
	ClientSecret string `json:"ClientSecret"`
	RedditUser   string `json:"RedditUser"`
	RedditPass   string `json:"RedditPass"`
}

// ReadConfig for json File
func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	ClientID = config.ClientID
	ClientSecret = config.ClientSecret
	RedditUser = config.RedditUser
	RedditPass = config.RedditPass

	return nil
}
