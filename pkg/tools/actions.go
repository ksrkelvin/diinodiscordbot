package tools

import (
	"diinoBot/pkg/config"
	"log"

	"github.com/bwmarrin/discordgo"
)

// Respond - Envia uma resposta para o usuário no Discord
func Respond(s *discordgo.Session, i *discordgo.InteractionCreate, content string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}

// RegisterCommands - Registra os Slash Commands no Discord
func RegisterCommands(s *discordgo.Session) {
	for _, cmd := range config.Commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			log.Fatalf("Erro ao registrar comando %s: %v", cmd.Name, err)
		}
		log.Printf("Comando registrado: /%s", cmd.Name)
	}
}
