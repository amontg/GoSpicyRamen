package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// struct to get tokens and keys
type configuration struct {
	BotPrefix  string
	BotToken   string
	YoutubeKey string
	BotID      string
	RedSec     string
	RedID      string
	RedUN      string
	RedPW      string
}

// config contains needed env variables
var config *configuration

// Load env variables
func Load() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file: ", err)
	}

	config = &configuration{
		BotPrefix:  os.Getenv("BOT_PREFIX"),
		BotToken:   os.Getenv("BOT_TOKEN"),
		YoutubeKey: os.Getenv("YOUTUBE_KEY"),
		BotID:      os.Getenv("BOT_ID"),
		RedSec:     os.Getenv("REDSEC"),
		RedID:      os.Getenv("REDID"),
		RedUN:      os.Getenv("REDUSER"),
		RedPW:      os.Getenv("REDPASS"),
	}

	//fmt.Println("Config Check: ", config.BotToken)
}

func GetBotPrefix() string {
	//fmt.Println(config.BotPrefix)
	return config.BotPrefix
}

func GetBotToken() string {
	return config.BotToken
}

func GetYoutubeKey() string {
	return config.YoutubeKey
}

func GetBotId() string {
	return config.BotID
}

func GetRedId() string {
	return config.RedID
}

func GetRedSec() string {
	return config.RedSec
}
