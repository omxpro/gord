package commands

import "github.com/omxpro/discordgo
"

type ClientState interface {
	GetSelectedGuild() *discordgo.Guild
	GetSelectedChannel() *discordgo.Channel
}
