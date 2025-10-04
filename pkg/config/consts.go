package config

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	BotToken = os.Getenv("TOKEN")
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "play",
			Description: "Reproduz áudio de um link do YouTube",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "url",
					Description: "URL do vídeo do YouTube",
					Required:    true,
				},
			},
		},
		{
			Name:        "stop",
			Description: "Para a reprodução e faz o bot sair do canal de voz",
		},
	}
)
