package bot

import (
	"fmt"

	"github.com/Anjasfedo/go-discord-ping-bot/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string

var GoBot *discordgo.Session

func Start() {
	GoBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := GoBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotID = user.ID

	GoBot.AddHandler(messageHandler)

	err = GoBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is Running")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}

	if m.Content == "anjas" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Anjas Gantenk")
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}
}
