package handlers

import (
	"fmt"

	paginator "github.com/TopiSenpai/dgo-paginator"
	"github.com/amontg/GoSpicyRamen/src/config"
	"github.com/amontg/GoSpicyRamen/src/context"

	//duck "github.com/amontg/GoSpicyRamen/src/duckduckgo"

	//"github.com/amontg/GoSpicyRamen/src/reddit"
	"github.com/amontg/GoSpicyRamen/src/utils"

	"github.com/amontg/GoSpicyRamen/src/youtube"

	"strings"

	"github.com/bwmarrin/discordgo"
)

// written by TopiSenpai
var manager = paginator.NewManager(paginator.WithButtonsConfig(paginator.ButtonsConfig{
	First: nil,
	Back: &paginator.ComponentOptions{
		Emoji: discordgo.ComponentEmoji{
			Name: "Left Arrow",
			ID:   "958424605339582584",
		},
		Style: discordgo.PrimaryButton,
	},
	Next: &paginator.ComponentOptions{
		Emoji: discordgo.ComponentEmoji{
			Name: "Right Arrow",
			ID:   "958424604882386966",
		},
		Style: discordgo.PrimaryButton,
	},
	Stop: &paginator.ComponentOptions{
		Emoji: discordgo.ComponentEmoji{
			Name: "Trash",
			ID:   "982065746635005982",
		},
		Style: discordgo.DangerButton,
	},
}))

// ---

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

		err := manager.CreateMessage(context.Dg, m.ChannelID, youtube.YtSearch(strings.Join(ytQuery, "%20"), m))
		if err != nil {
			fmt.Println(err)
		}
	case prefix + "reddit":
		utils.SimpleMessage(m.ChannelID, "https://c.tenor.com/X8q1Q4i3qmwAAAAC/nervous-glance.gif")
		//rQuery := utils.PopDown(cmd)
		//reddit.RedditSearch(strings.Join(rQuery, "%20"), m)
	case prefix + "find":
		//fQuery := utils.PopDown(cmd)

		// err := manager.CreateMessage(context.Dg, m.ChannelID, duck.SearchThis(strings.Join(fQuery, "%20")))
		// if err != nil {
		// 	fmt.Println(err)
		// }
	default:
		return
	}
}

func InteractionCreateHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	// first get the message data

}
