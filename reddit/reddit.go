package reddit

import (
	"fmt"
	"log"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/jmrlgg/godiscord/config"
	"github.com/jzelinskie/geddit"
)

var (
	RedditParms string // RedditParms Exported from Bot File
)

// Search for movies tvshows
func Search(RedditParms string, s *discordgo.Session, m *discordgo.MessageCreate) {
	o, err := geddit.NewOAuthSession(
		config.ClientID,
		config.ClientSecret,
		"Testing Geddit Bot with OAuth v0.1 by u/tehmass ",
		"",
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("test")
	// Login using our personal reddit account.
	err = o.LoginAuth(config.RedditUser, config.RedditPass)
	if err != nil {
		log.Fatal(err)
	}

	// We can pass options to the query if desired (blank for now).
	opts := geddit.ListingOptions{}
	// Fetch posts from r/videos, sorted by Hot.
	posts, err := o.SubredditSubmissions("bestofstreamingvideo", geddit.NewSubmissions, opts)
	if err != nil {
		log.Fatal(err)
	}

	// Print the title and URL of each post that has a youtube.com domain.
	for _, p := range posts {
		if p.Title != "[Request]" {
			if regexp.MustCompile(RedditParms).MatchString(p.Title) == true {
				fmt.Printf("%s (%d) - %s\n", p.Title, p.Score, p.URL)
				_, _ = s.ChannelMessageSend(m.ChannelID, p.Title+"\n"+p.URL)
			}

		}
	}
	return
}
