package commands

import "github.com/gord-project/discordgo"

type ClientState interface {
	GetSelectedGuild() *discordgo.Guild
	GetSelectedChannel() *discordgo.Channel
}
