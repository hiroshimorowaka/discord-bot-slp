package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
}
