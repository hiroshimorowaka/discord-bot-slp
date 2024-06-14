package status

import (
	"go_bot/bot/handlers/status/embeds"
	"go_bot/lib"

	"github.com/bwmarrin/discordgo"
)

func Run(s *discordgo.Session, i *discordgo.InteractionCreate) {

	response, err := lib.GetServerInfo()

	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: embeds.ErrorEmbed(err),
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embeds.SuccessEmbed(response),
		},
	})
}
