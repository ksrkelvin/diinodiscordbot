package audio

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/ogg"
)

func Play(session *discordgo.Session, guildID, channelID, audioURL string) error {
	voice, err := session.ChannelVoiceJoin(guildID, channelID, false, false)
	if err != nil {
		return fmt.Errorf("failed to join voice channel: %w", err)
	}

	dlp := exec.Command(
		"yt-dlp",
		"--extract-audio", audioURL,
		"-o", "-", "--playlist-items", "1",
	)
	dlpPipe, err := dlp.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get yt-dlp pipe: %w", err)
	}
	dlp.Stderr = os.Stderr

	ffmpeg := exec.Command(
		"ffmpeg",
		"-i", "-",
		"-f", "opus",
		"-frame_duration", "20",
		"-ar", "48000",
		"-ac", "2",
		"-",
	)
	ffmpegPipe, err := ffmpeg.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get ffmpeg pipe: %w", err)
	}
	ffmpeg.Stdin = dlpPipe
	ffmpeg.Stderr = os.Stderr

	if err := dlp.Start(); err != nil {
		return err
	}
	if err := ffmpeg.Start(); err != nil {
		return err
	}

	pageDecoder := ogg.NewDecoder(ffmpegPipe)
	pageDecoder.Decode()

	voice.Speaking(true)
	packetDecoder := ogg.NewPacketDecoder(pageDecoder)
	for {
		packet, _, err := packetDecoder.Decode()
		if err != nil {
			return err
		}

		voice.OpusSend <- packet
	}
}
