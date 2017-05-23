package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/discord-fortune-bot/fortune"
)

const helpMessage = `
Use "/fortune" to get me to print out a Unix fortune.
Use "/fortune offensive" if you want me to be filthy.

For more information, check my Github repo: https://github.com/fsufitch/discord-fortune-bot.
`

func runBotAsync(token string, stop <-chan bool, doneChan chan<- error) {
	log.Printf("Starting bot with token %s", token)
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		doneChan <- err
		return
	}

	dg.AddHandler(handleMessage)

	err = dg.Open()
	defer dg.Close()
	if err != nil {
		doneChan <- err
		return
	}

	<-stop
	log.Print("Stopping bot...")
	doneChan <- nil
	return
}

// RunBot runs the Discord bot and returns a channel that
// puts out an error value (or nil) when the bot terminates
func RunBot(token string, stop <-chan bool) <-chan error {
	doneChan := make(chan error)
	go runBotAsync(token, stop, doneChan)

	return doneChan
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.ToLower(m.Content)
	fields := strings.Fields(content)
	if len(fields) < 1 || fields[0] != "/fortune" {
		return
	}

	offensive := false

	if len(fields) >= 2 {
		if fields[1] == "help" {
			s.ChannelMessageSend(m.ChannelID, helpMessage)
			return
		}
		if fields[1] == "offensive" {
			offensive = true
		} else {
			msg := fmt.Sprintf("Unknown fortune action: %s", fields[1])
			s.ChannelMessageSend(m.ChannelID, msg)
			return
		}
	}

	f, err := fortune.GetFortune(offensive)
	if err == nil {
		s.ChannelMessageSend(m.ChannelID, f)
	} else {
		s.ChannelMessageSend(m.ChannelID, "Error getting fortune: "+err.Error())
	}
}
