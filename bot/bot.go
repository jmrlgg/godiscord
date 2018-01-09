package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jmrlgg/godiscord/config"
	"github.com/jmrlgg/godiscord/reddit"
)

// const token string = "Mzg2MjYzNTUxNTA2OTA3MTQ5.DQNYFw.nloki1PSCCF-a9m74iQM8fGdcGg"

// BotID Public Comment
var BotID string

var goBot *discordgo.Session

// Start Bot Start Function
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

	fmt.Println("Bot is Up!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// if strings.HasPrefix(m.Content, config.BotPrefix)
	if m.Content[:1] == config.BotPrefix {
		if m.Author.ID == BotID {
			return
		}

		if m.Content[1:] == "tv" {
			fmt.Println("Owner COmmand executred")
			_, _ = s.ChannelMessageSend(m.ChannelID, "tehmass : @jmrlgg "+m.Content[1:])

		}

		if strings.Contains(m.Content, config.BotPrefix+"tv ") {
			RedditParms := m.Content[4:]
			fmt.Println(m.Content[4:])
			//_, _ = s.ChannelMessageSend(m.ChannelID, m.Content[4:])
			reddit.Search(RedditParms, s, m)
		}
	}
}
