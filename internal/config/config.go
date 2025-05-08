package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken      string
	ReactionEmoji string
	MessageRegex  string
	ClientID      string
	GuildID       string
	Env           string
}

const (
	BotTokenEnvVar = "DISCORD_BOT_TOKEN"
	CLientIDEnvVar = "DISCORD_CLIENT_ID"
	GuildIDEnvVar  = "DISCORD_GUILD_ID"
	EnvEnvVar      = "ENV"

	DefaultReactionEmoji = "✅"
	DefaultMessageRegex  = "toto"
)

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf(".env file must be defined")
	}

	botToken := os.Getenv(BotTokenEnvVar)
	if botToken == "" {
		return nil, fmt.Errorf("env variable %s must be defined", BotTokenEnvVar)
	}

	clientID := os.Getenv(CLientIDEnvVar)
	if botToken == "" {
		return nil, fmt.Errorf("client id %s must be defined", GuildIDEnvVar)
	}

	guildID := os.Getenv(GuildIDEnvVar)
	if botToken == "" {
		return nil, fmt.Errorf("guild id %s must be defined", GuildIDEnvVar)
	}

	env := os.Getenv(EnvEnvVar)
	if env != "dev" && env != "prod" {
		return nil, fmt.Errorf("env %s must be either \"dev\" or \"prod\"", EnvEnvVar)
	}

	return &Config{
		BotToken:      botToken,
		ClientID:      clientID,
		GuildID:       guildID,
		ReactionEmoji: DefaultReactionEmoji,
		MessageRegex:  DefaultMessageRegex,
		Env:           env,
	}, nil
}
