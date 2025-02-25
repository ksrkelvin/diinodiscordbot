package router

import (
	"diinoBot/pkg/audio"
	"diinoBot/pkg/tools"
	"log"

	"github.com/bwmarrin/discordgo"
)

// InteractionCreate - Lida com as interações (respostas aos comandos do usuário)
func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case "play":
		url := i.ApplicationCommandData().Options[0].StringValue()
		guildID := i.GuildID

		// Verifica se o bot está em um canal antes de sair
		if tools.IsBotInVoiceChannel(s, guildID) {
			if err := tools.LeaveVoiceChannel(s, guildID); err != nil {
				log.Println("Error leaving voice channel:", err)
				tools.Respond(s, i, "Erro ao sair do canal de voz.")
				return
			}
		}

		// Obtém o canal de voz do usuário que enviou o comando
		voiceChannelID, err := tools.GetVoiceChannel(s, guildID, i.Member.User.ID)
		if err != nil {
			log.Println("Error getting voice channel:", err)
			tools.Respond(s, i, "Você precisa estar em um canal de voz para usar esse comando.")
			return
		}

		// Responde ao usuário que a música vai tocar
		tools.Respond(s, i, "🎶 Tocando: "+url)

		// Toca o áudio
		if err := audio.Play(s, guildID, voiceChannelID, url); err != nil {
			log.Println("Error playing audio:", err)
			tools.Respond(s, i, "Erro ao tocar o áudio.")
		}

	case "stop":
		guildID := i.GuildID
		// Verifica se o bot está em um canal antes de sair
		if !tools.IsBotInVoiceChannel(s, guildID) {
			tools.Respond(s, i, "O bot não está em nenhum canal de voz.")
			return
		}

		if err := tools.LeaveVoiceChannel(s, guildID); err != nil {
			log.Println("Error leaving voice channel:", err)
			tools.Respond(s, i, "Erro ao sair do canal de voz.")
			return
		}
		tools.Respond(s, i, "⏹ Bot saiu do canal de voz.")
	}
}
