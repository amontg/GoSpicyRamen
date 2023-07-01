package handlers

import (
	"fmt"
	"log"

	paginator "github.com/TopiSenpai/dgo-paginator"
	"github.com/amontg/GoSpicyRamen/src/config"
	"github.com/amontg/GoSpicyRamen/src/context"

	//duck "github.com/amontg/GoSpicyRamen/src/duckduckgo"

	"github.com/amontg/GoSpicyRamen/src/reddit"
	"github.com/amontg/GoSpicyRamen/src/urbandictionary"
	"github.com/amontg/GoSpicyRamen/src/utils"

	"github.com/amontg/GoSpicyRamen/src/youtube"

	"strings"

	"github.com/bwmarrin/discordgo"
)

// written by @TopiSenpai
var manager = paginator.NewManager(
	paginator.WithButtonsConfig(paginator.ButtonsConfig{
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
				Name: "Stop",
				ID:   "1111762026453278822",
			},
			Style: discordgo.DangerButton,
		},
		First: &paginator.ComponentOptions{
			Emoji: discordgo.ComponentEmoji{
				Name: "First",
				ID:   "1111760730606293144",
			},
			Style: discordgo.PrimaryButton,
		},
		Last: &paginator.ComponentOptions{
			Emoji: discordgo.ComponentEmoji{
				Name: "Last",
				ID:   "1111760732451766272",
			},
			Style: discordgo.PrimaryButton,
		},
	}),
	paginator.WithNotYourPaginatorMessage("This paginator is not yours."),
)

// ---

func AddHandlers() {
	// this is used to register handlers for events
	// context.dg.AddHandler(ReadyHandler)
	context.Dg.AddHandler(manager.OnInteractionCreate)
	context.Dg.AddHandler(MessageCreateHandler)
}

// Create(From)Interaction: The response based on an interaction
// OnInteractionCreate: Event handler

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
	query := strings.Join(utils.PopDown(cmd), "%20")

	switch cmd[0] {
	case prefix + "ping":
		utils.SimpleMessage(m.ChannelID, "Please shut up.")
	case prefix + "yt":
		result := youtube.YtSearch(query, m)
		err := manager.CreateMessage(context.Dg, m.ChannelID, result)
		if err != nil {
			fmt.Println(err)
			log.Panic(err)
		}
	case prefix + "reddit":
		//utils.SimpleMessage(m.ChannelID, "https://c.tenor.com/X8q1Q4i3qmwAAAAC/nervous-glance.gif")
		result := reddit.RedditSearch(query, m)
		if result != nil {
			err := manager.CreateMessage(context.Dg, m.ChannelID, result)
			if err != nil {
				fmt.Println(err)
				log.Panic(err)
			}
		}
	case prefix + "ud":
		result := urbandictionary.UrbanDictSearch(query, m)
		if result != nil {
			err := manager.CreateMessage(context.Dg, m.ChannelID, result)
			if err != nil {
				fmt.Println(err)
				log.Panic(err)
			}
		}
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
