package handlers

import (
	"github.com/amontg/GoSpicyRamen/src/config"
	"github.com/amontg/GoSpicyRamen/src/context"
	"github.com/amontg/GoSpicyRamen/src/utils"

	"github.com/amontg/GoSpicyRamen/src/youtube"

	"strings"

	//"fmt"

	"github.com/bwmarrin/discordgo"
)

func AddHandlers() {
	// this is used to register handlers for events
	// context.dg.AddHandler(ReadyHandler)
	context.Dg.AddHandler(MessageCreateHandler)

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
		utils.SendChannelMessage(m.ChannelID, "Please shut up.")
	case prefix + "yt":
		// we want to use everything besides cmd[0] for our search
		ytQuery := utils.PopDown(cmd)
		//fmt.Println(strings.Join(ytQuery, "%20")) -- youtube search puts a %20 instead of spaces
		youtube.YtSearch(strings.Join(ytQuery, "%20"), m)
		// func YtSearch(query string, channelID string)
		// utils.SendChannelMessage(m.ChannelID, youtube.YtSearch(cmd[1]))
	default:
		return
	}
}
