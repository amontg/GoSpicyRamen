package utils

import (
	// "os"
	"fmt"

	paginator "github.com/TopiSenpai/dgo-paginator"
	"github.com/amontg/GoSpicyRamen/src/context"
	"github.com/bwmarrin/discordgo"
)

// sends shit to the session thank you

func SimpleMessage(channelID string, message string) { // Sends a message to the channel. channelID string, message string)
	context.Dg.ChannelMessageSend(channelID, message)
}

func ComplexMessage(channelID string, msg *discordgo.MessageSend) { // Sends a message to the channel but with the buttons: (back) (forward) (random) (kill)
	_, err := context.Dg.ChannelMessageSendComplex(channelID, msg)
	if err != nil {
		fmt.Println(err)
	}
}

func EmptyComplex() *discordgo.MessageSend {
	var m *discordgo.MessageSend = new(discordgo.MessageSend)

	m.Content = "Well, uh... There's an error."

	return m
}

func EmptyPaginator() *paginator.Paginator {
	var m *paginator.Paginator = new(paginator.Paginator)

	return m
}

func EditComlex(m *discordgo.MessageEdit) {
	_, err := context.Dg.ChannelMessageEditComplex(m)
	if err != nil {
		fmt.Println(err)
	}
}

// function to remove the first element of a slice and shift everything down
func PopDown(slice []string) []string {
	slice = slice[1:]
	return slice
}
