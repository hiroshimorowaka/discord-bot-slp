package embeds

import (
	"github.com/bwmarrin/discordgo"
)

func ErrorEmbed(err error) []*discordgo.MessageEmbed {
	return []*discordgo.MessageEmbed{{
		Title: "Server Status",
		Color: 10038562,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Status",
				Value:  "Probably Offline",
				Inline: false,
			},
			{
				Name:   "Informations",
				Value:  err.Error(),
				Inline: false,
			},
		},
		Type: discordgo.EmbedTypeRich,
	}}

}
