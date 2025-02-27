package tools

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// GetVoiceChannel - Função para obter o canal de voz de um usuário
func GetVoiceChannel(s *discordgo.Session, guildID string, userID string) (string, error) {
	fmt.Println("Fetching voice channel for Guild ID:", guildID) // Debug

	// Verifica se o bot está conectado à guilda no cache
	guild, err := s.State.Guild(guildID)
	if err != nil {
		fmt.Println("Guild not found in cache, fetching from API...")

		// Se não estiver no cache, tenta buscar a guilda pela API
		guild, err = s.Guild(guildID)
		if err != nil {
			return "", fmt.Errorf("failed to get guild: %w", err)
		}
	}

	// Procura o estado de voz do usuário na guilda
	for _, vs := range guild.VoiceStates {
		if vs.UserID == userID {
			return vs.ChannelID, nil
		}
	}

	return "", fmt.Errorf("user is not in a voice channel")
}

// LeaveVoiceChannel -  Função para fazer o bot sair do canal de voz
func LeaveVoiceChannel(s *discordgo.Session, guildID string) (err error) {
	// Obtém os canais da guilda
	channels, err := s.GuildChannels(guildID)
	if err != nil {
		return fmt.Errorf("failed to get guild channels: %w", err)
	}

	for _, channel := range channels {
		if channel.Type == discordgo.ChannelTypeGuildVoice {
			// Encontra a conexão de voz ativa do bot
			voiceConnection := s.VoiceConnections[guildID]

			// Faz o bot sair do canal de voz
			voiceConnection.Disconnect()
			log.Println("Bot left the voice channel:", channel.ID)
			return nil
		}
	}

	return fmt.Errorf("no voice channel found to leave in guild: %s", guildID)
}
