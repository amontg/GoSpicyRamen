package main

import (
	//"encoding/json"
	"flag"
	"fmt"

	//"io/ioutil"
	//"net/http"
	"os"
	"os/signal"

	//"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string // initialized to be used for the token
)

func init() {
	// create command-line flags. returns pointer
	// flag.StringVar(&flagvar, "identifier", "default output", "help message")
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// create a new discord session
	// err = error type, := is a variable declaration
	// discordgo.New constructs a new Discord client
	dg, err := discordgo.New("Bot " + Token)

	if err != nil { // if error exists, print it
		fmt.Println("Error creating session: ", err)
		return
	}

	// register messageCreate func as callback for events
	dg.AddHandler(messageCreate)

	// only care about receiving message events, for now
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// open websocket connection and start listening
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	// wait til closed with CTRL-C/termination signal received
	fmt.Println("The bot is running, exit with CTRL-C.")
	sc := make(chan os.Signal, 1) // a channel to get and handle Unix signals
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Close session
	dg.Close()
}

// function to be called when a new message is created (thanks, addHandler)
// func name(var type, var type) IE discordgo.Session bot struct, discordgo.MessageCreate struct
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// ignore the bot's messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == ".ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Please shut up.")
	}
}

/*
	1. Export token
		$ export BOT_TOKEN=<token>
	2. Locally run bot
		$ go run <file> -t $BOT_TOKEN
*/
