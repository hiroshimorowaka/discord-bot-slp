package embeds

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/sch8ill/mclib/slp"
)

func SuccessEmbed(response *slp.Response) []*discordgo.MessageEmbed {

	var players = ""
	for _, player := range response.Players.Sample {
		players = players + player.Name + "\n"
	}

	return []*discordgo.MessageEmbed{{
		Title: "Server Status",
		Color: 5763719,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Status",
				Value:  "Online",
				Inline: true,
			},
			{
				Name:   "Ping",
				Value:  fmt.Sprintf("%dms", response.Latency),
				Inline: true,
			},
			{
				Name:   "Player Count",
				Value:  fmt.Sprintf("%d/%d", response.Players.Online, response.Players.Max),
				Inline: true,
			},
			{
				Name:   "Players",
				Value:  players,
				Inline: false,
			},
			{
				Name:   "Server IP",
				Value:  os.Getenv("HOST"),
				Inline: true,
			},
			{
				Name:   "Version",
				Value:  response.Version.Name,
				Inline: true,
			},
		},
		Type: discordgo.EmbedTypeRich,
	}}

}
