package discordavatarplugin

import (
	"regexp"
	"strings"

	"github.com/dwatson251/bruxism"
	"github.com/dwatson251/discordgo"
)

var userIDRegex = regexp.MustCompile("<@!?([0-9]*)>")

func avatarMessageFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message) {
	if service.Name() == bruxism.DiscordServiceName && !service.IsMe(message) {
		if bruxism.MatchesCommand(service, "avatar", message) {
			query := strings.Join(strings.Split(message.RawMessage(), " ")[1:], " ")

			id := message.UserID()
			match := userIDRegex.FindStringSubmatch(query)
			if match != nil {
				id = match[1]
			}

			discord := service.(*bruxism.Discord)

			u, err := discord.Session.User(id)
			if err != nil {
				return
			}

			service.SendMessage(message.Channel(), discordgo.EndpointUserAvatar(u.ID, u.Avatar))
		}
	}
}

func avatarHelpFunc(bot *bruxism.Bot, service bruxism.Service, message bruxism.Message, detailed bool) []string {
	if detailed {
		return nil
	}
	return bruxism.CommandHelp(service, "avatar", "[@username]", "Returns a big version of your avatar, or a users avatar if provided.")
}

// New creates a new discordavatar plugin.
func New() bruxism.Plugin {
	p := bruxism.NewSimplePlugin("discordavatar")
	p.MessageFunc = avatarMessageFunc
	p.HelpFunc = avatarHelpFunc
	return p
}
