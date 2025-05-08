package bot

import (
	"context"
	"log"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/snowflake/v2"

	"github.com/sebferrer/discord-bot-go-template/internal/bot/commands"
	"github.com/sebferrer/discord-bot-go-template/internal/bot/eventhandler"
	"github.com/sebferrer/discord-bot-go-template/internal/config"
)

type BotClient struct {
	bot.Client
}

func NewClient(cfg *config.Config) (*BotClient, error) {
	handler := eventhandler.New(cfg)

	client, err := disgo.New(cfg.BotToken,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuildMessages|
					gateway.IntentGuilds|
					gateway.IntentMessageContent,
			),
		),
		bot.WithEventListeners(handler),
	)
	if err != nil {
		return nil, err
	}

	handler.InjectBotClient(client)

	// Register slash commands if GuildID is configured
	if cfg.Env == "dev" && cfg.GuildID != "" {
		guildID := snowflake.MustParse(cfg.GuildID)
		log.Printf("Registering commands for guild %s (dev mode)", guildID)
		commands.RegisterCommandsForGuild(client, guildID)
	} else {
		log.Println("Registering global commands (prod mode)")
		commands.RegisterGlobalCommands(client)
	}

	return &BotClient{client}, nil
}

func (b *BotClient) Connect(ctx context.Context) error {
	return b.Client.OpenGateway(ctx)
}

func (b *BotClient) Close(ctx context.Context) {
	b.Client.Close(ctx)
}

func (b *BotClient) ApplicationID() snowflake.ID {
	return b.Client.ApplicationID()
}
