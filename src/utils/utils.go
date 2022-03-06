package utils

import (
	// "os"
	// "fmt"

	"github.com/amontg/GoSpicyRamen/src/context"
)

func SendChannelMessage(channelID string, message string) { // Sends a message to the channel. channelID string, message string)
	context.Dg.ChannelMessageSend(channelID, message)
}

// function to remove the first element of a slice and shift everything down
func PopDown(slice []string) []string {
	slice = slice[1:]
	return slice
}
