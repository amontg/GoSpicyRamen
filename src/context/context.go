package context

import (
	"fmt"

	"github.com/amontg/GoSpicyRamen/src/config"
	"github.com/bwmarrin/discordgo"
)

var Dg *discordgo.Session

func Initialize(botToken string) {
	var err error
	// initialize the new discord session
	Dg, err = discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println("Error starting a session: ", err)
		return
	}
	//fmt.Println("Init Check: ", botToken)
	var inviteLink string = "https://discord.com/api/oauth2/authorize?client_id="
	botId := config.GetBotId()
	fmt.Println("Invite me: ", inviteLink, botId, "&permissions=8&scope=bot")
}

func OpenConnection() {
	// create the connection
	if err := Dg.Open(); err != nil {
		fmt.Println("Error creating a connection: ", err)
		return
	}
}
