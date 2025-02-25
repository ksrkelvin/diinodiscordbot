package tools

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// GetVoiceChannel - Função para obter o ID do canal de voz de uma guilda
func GetVoiceChannel(s *discordgo.Session, guildID string, userID string) (string, error) {
	guild, err := s.State.Guild(guildID)
	if err != nil {
		return "", fmt.Errorf("failed to get guild: %w", err)
	}

	for _, vs := range guild.VoiceStates {
		if vs.UserID == userID {
			return vs.ChannelID, nil
		}
	}

	return "", fmt.Errorf("user is not in a voice channel")
}

// IsBotInVoiceChannel - Função para verificar se o bot está em um canal de voz
func IsBotInVoiceChannel(s *discordgo.Session, guildID string) bool {
	voiceState, err := s.State.VoiceState(guildID, s.State.User.ID)
	return err == nil && voiceState != nil
}

// LeaveVoiceChannel -  Função para fazer o bot sair do canal de voz
func LeaveVoiceChannel(s *discordgo.Session, guildID string) error {
	// Obtém os canais da guilda
	channels, err := s.GuildChannels(guildID)
	if err != nil {
		return fmt.Errorf("failed to get guild channels: %w", err)
	}

	// Encontra o canal de voz que o bot está
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
