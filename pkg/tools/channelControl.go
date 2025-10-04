package tools

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func GetVoiceChannel(s *discordgo.Session, guildID string, userID string) (string, error) {
	fmt.Println("Fetching voice channel for Guild ID:", guildID)

	guild, err := s.State.Guild(guildID)
	if err != nil {
		fmt.Println("Guild not found in cache, fetching from API...")

		guild, err = s.Guild(guildID)
		if err != nil {
			return "", fmt.Errorf("failed to get guild: %w", err)
		}
	}

	for _, vs := range guild.VoiceStates {
		if vs.UserID == userID {
			return vs.ChannelID, nil
		}
	}

	return "", fmt.Errorf("user is not in a voice channel")
}

func LeaveVoiceChannel(s *discordgo.Session, guildID string) (err error) {
	channels, err := s.GuildChannels(guildID)
	if err != nil {
		return fmt.Errorf("failed to get guild channels: %w", err)
	}

	for _, channel := range channels {
		if channel.Type == discordgo.ChannelTypeGuildVoice {
			voiceConnection := s.VoiceConnections[guildID]

			voiceConnection.Disconnect()
			log.Println("Bot left the voice channel:", channel.ID)
			return nil
		}
	}

	return fmt.Errorf("no voice channel found to leave in guild: %s", guildID)
}
