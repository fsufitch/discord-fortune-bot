package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/fsufitch/discord-fortune-bot/bot"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()

	if Token == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	botStop := make(chan bool)
	botDone := bot.RunBot(Token, botStop)
	fmt.Println("Bot started")

	intSignal := make(chan os.Signal, 1)
	signal.Notify(intSignal, os.Interrupt)

	for {
		select {
		case <-intSignal:
			botStop <- true
		case err := <-botDone:
			if err != nil {
				fmt.Println("Had bot error:", err)
			}
			return
		}
	}
}
