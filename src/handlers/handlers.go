package handlers

import (
	"fmt"

	paginator "github.com/TopiSenpai/dgo-paginator"
	"github.com/amontg/GoSpicyRamen/src/config"
	"github.com/amontg/GoSpicyRamen/src/context"
	"github.com/amontg/GoSpicyRamen/src/utils"

	"github.com/amontg/GoSpicyRamen/src/youtube"

	"strings"

	"github.com/bwmarrin/discordgo"
)

var manager = paginator.NewManager()

func AddHandlers() {
	// this is used to register handlers for events
	// context.dg.AddHandler(ReadyHandler)
	context.Dg.AddHandler(MessageCreateHandler)
	//context.Dg.AddHandler(InteractionCreateHandler)

	context.Dg.AddHandler(manager.OnInteractionCreate)
}

func ReadyHandler(s *discordgo.Session, event *discordgo.GuildCreate) {

	// set playing status
	s.UpdateGameStatus(0, "Being a real bird.")
}

func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// prevent bot from answering itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	prefix := config.GetBotPrefix()
	// split up the given command
	cmd := strings.Split(m.Content, " ")

	switch cmd[0] {
	case prefix + "ping":
		utils.SimpleMessage(m.ChannelID, "Please shut up.")
	case prefix + "yt":
		// func YtSearch(query string, channelID string)
		ytQuery := utils.PopDown(cmd)

		// this func will send the message using paginator's CreateMessage()
		err := manager.CreateMessage(context.Dg, m.ChannelID, youtube.YtSearch(strings.Join(ytQuery, "%20"), m))
		if err != nil {
			fmt.Println(err)
		}
	default:
		return
	}
}

func InteractionCreateHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	// first get the message data

}
