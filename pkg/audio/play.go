package audio

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/ogg"
)

// Play -  Função para executar o áudio no canal de voz
func Play(session *discordgo.Session, guildID, channelID, audioURL string) error {
	voice, err := session.ChannelVoiceJoin(guildID, channelID, false, false)
	if err != nil {
		return fmt.Errorf("failed to join voice channel: %w", err)
	}

	// Configurando fonte de áudio
	dlp := exec.Command(
		"yt-dlp",
		"--extract-audio", audioURL, // Agora o link é passado como argumento
		"-o", "-", "--playlist-items", "1",
	)
	dlpPipe, err := dlp.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get yt-dlp pipe: %w", err)
	}
	dlp.Stderr = os.Stderr // Obter saída informativa do yt-dlp

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
	ffmpeg.Stdin = dlpPipe // Associa saída do yt-dlp à entrada do ffmpeg
	ffmpeg.Stderr = os.Stderr

	if err := dlp.Start(); err != nil {
		return err
	}
	if err := ffmpeg.Start(); err != nil {
		return err
	}

	pageDecoder := ogg.NewDecoder(ffmpegPipe)
	pageDecoder.Decode()

	// Sinaliza ao Discord que estamos enviando áudio
	voice.Speaking(true)
	packetDecoder := ogg.NewPacketDecoder(pageDecoder)
	for {
		packet, _, err := packetDecoder.Decode()
		if err != nil {
			// Esperamos que o único erro seja io.EOF
			return err
		}

		// Enviando o áudio para o Discord
		voice.OpusSend <- packet
	}
}
