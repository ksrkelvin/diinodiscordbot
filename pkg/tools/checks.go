package tools

import "github.com/bwmarrin/discordgo"

func IsBotInVoiceChannel(s *discordgo.Session, guildID string) bool {
	voiceState, err := s.State.VoiceState(guildID, s.State.User.ID)
	return err == nil && voiceState != nil
}
