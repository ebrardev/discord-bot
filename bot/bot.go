package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ebrardev/discord-bot/config"
)

var BotID string
var goBot *discordgo.Session

func Start() {

	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())

	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)
	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("deneme 1-2,Bot is running...,")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}

	// Ignore all messages created by the bot itself
	if m.Content == "php" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "henüz ölmedi")

	}
	if m.Content == "gel bir sarayım" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "aşkın olayım")

	}
	if m.Content == "kurtlar vadisi vs ezel ?" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "ezel tabi ki")

	}
	if m.Content == "faze" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "çöp takım")

	}

	if m.Content == "React vs vue ?" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Vueeee")

	}
	if m.Content == "fenerbahce icin" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Ali koç istifa")

	}

}
