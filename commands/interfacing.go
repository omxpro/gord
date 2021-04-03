package commands

import "github.com/cainy-a/discordgo"

type ClientState interface {
	GetSelectedGuild() *discordgo.Guild
	GetSelectedChannel() *discordgo.Channel
}
