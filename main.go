package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/amontg/GoSpicyRamen/src/bot"
	"github.com/amontg/GoSpicyRamen/src/config"
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
	config.Load()
	bot.Start()

	fmt.Println("Bot is running, exit with CTRL+C")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Stop()
}
