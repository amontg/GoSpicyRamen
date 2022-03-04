package utils

import (
	// "os"
	// "fmt"

	"github.com/amontg/GoSpicyRamen/src/context"
)

// send messages
func SendChannelMessage(channelID string, message string) {
	context.Dg.ChannelMessageSend(channelID, message)
}
