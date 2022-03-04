package context

import (
	"fmt"

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
	fmt.Println("Init Check: ", botToken)
}

func OpenConnection() {
	// create the connection
	if err := Dg.Open(); err != nil {
		fmt.Println("Error creating a connection: ", err)
		return
	}
}
