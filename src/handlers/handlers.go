package handlers

import (
	paginator "github.com/TopiSenpai/dgo-paginator"
	"github.com/amontg/GoSpicyRamen/src/config"
	"github.com/amontg/GoSpicyRamen/src/context"
	"github.com/amontg/GoSpicyRamen/src/utils"

	"github.com/amontg/GoSpicyRamen/src/youtube"

	"strings"

	"github.com/bwmarrin/discordgo"
)

// written by TopiSenpai
var manager = paginator.NewManager(paginator.WithButtonsConfig(paginator.ButtonsConfig{
	First: &paginator.ComponentOptions{
		Emoji: discordgo.ComponentEmoji{
			Name: "⏮",
		},
		Style: discordgo.PrimaryButton,
	},
	Back: &paginator.ComponentOptions{
		Emoji: discordgo.ComponentEmoji{
			Name: "◀",
		},
		Style: discordgo.PrimaryButton,
	},
	Stop: &paginator.ComponentOptions{
		Emoji: discordgo.ComponentEmoji{
			Name: "🗑",
		},
		Style: discordgo.DangerButton,
	},
	Next: &paginator.ComponentOptions{
		Emoji: discordgo.ComponentEmoji{
			Name: "▶",
		},
		Style: discordgo.PrimaryButton,
	},
	Last: &paginator.ComponentOptions{
		Emoji: discordgo.ComponentEmoji{
			Name: "⏩",
		},
		Style: discordgo.PrimaryButton,
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

		// this func will send the message using paginator's CreateMessage()
		/*
			err := manager.CreateMessage(context.Dg, m.ChannelID, )
			if err != nil {
				fmt.Println(err)
			}
		*/

		youtube.YtSearch(strings.Join(ytQuery, "%20"), m)

	default:
		return
	}
}

func InteractionCreateHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	// first get the message data

}
