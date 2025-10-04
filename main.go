package main

import (
	"diinoBot/pkg/config"
	"diinoBot/pkg/router"
	"diinoBot/pkg/tools"
	"log"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	session, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		log.Fatalln("failed to create bot session: " + err.Error())
	}

	// Registrando o manipulador de eventos para mensagens
	session.AddHandler(router.InteractionCreate)

	slog.Info("Opening bot session")
	if err := session.Open(); err != nil {
		log.Fatalln("failed to open bot session: " + err.Error())
	}
	defer session.Close()

	// Registra os comandos
	tools.RegisterCommands(session)

	// Mantém o bot rodando
	select {}
}
