package tools

import (
	"diinoBot/pkg/config"
	"log"

	"github.com/bwmarrin/discordgo"
)

func Respond(s *discordgo.Session, i *discordgo.InteractionCreate, content string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}

func RegisterCommands(s *discordgo.Session) {
	for _, cmd := range config.Commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			log.Fatalf("Erro ao registrar comando %s: %v", cmd.Name, err)
		}
		log.Printf("Comando registrado: /%s", cmd.Name)
	}
}
