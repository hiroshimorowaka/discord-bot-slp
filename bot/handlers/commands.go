package handlers

import (
	"go_bot/bot/handlers/status"

	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name: "status",
			// All commands and options must have a description
			// Commands/options without description will fail the registration
			// of the command.
			Description: "Ver o status atual do server do mine.",
		},
	}

	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"status": status.Run,
	}
)
