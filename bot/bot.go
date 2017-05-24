package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/fsufitch/discord-fortune-bot/fortune"
)

const errorTemplate = "```Error: %s\n\nTry \"/fortune -h\" to get help.```"

const messageTemplate = "```\n%s\n```"

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

	fields := strings.Fields(m.Content)
	if len(fields) < 1 || strings.ToLower(fields[0]) != "/fortune" {
		return
	}

	options, err := parseFlags(fields[1:])

	if options.TextOverride != "" {
		message := fmt.Sprintf(messageTemplate, options.TextOverride)
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	if err != nil {
		message := fmt.Sprintf(errorTemplate, err.Error())
		s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	f, err := fortune.GetFortune(options.Offensive, options.Length, options.Passthrough)
	if err == nil {
		message := fmt.Sprintf(messageTemplate, f)
		s.ChannelMessageSend(m.ChannelID, message)
	} else {
		message := fmt.Sprintf(errorTemplate, err.Error())
		s.ChannelMessageSend(m.ChannelID, message)
	}
}
