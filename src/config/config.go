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
	}

	fmt.Println("Config Check: ", config.BotToken)
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
