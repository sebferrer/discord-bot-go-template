package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sebferrer/discord-bot-go-template/internal/bot"
	"github.com/sebferrer/discord-bot-go-template/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	botClient, err := bot.NewClient(cfg)
	if err != nil {
		log.Fatalf("client creation: %v", err)
	}

	err = botClient.Connect(context.Background())
	if err != nil {
		log.Fatalf("gateway connexion: %v", err)
	}
	defer botClient.Close(context.Background())

	fmt.Printf("invitation link : https://discord.com/oauth2/authorize?client_id=%s&permissions=563364418144320&integration_type=0&scope=bot+applications.commands\n", cfg.ClientID)

	log.Printf("bot %s started. Ctrl+C to quit.\n", botClient.ApplicationID())

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	fmt.Println("Stopping bot...")
}
