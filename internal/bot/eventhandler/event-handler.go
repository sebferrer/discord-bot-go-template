package eventhandler

import (
	"fmt"
	"log"
	"regexp"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
	"github.com/sebferrer/discord-bot-go-template/internal/bot/commands"
	"github.com/sebferrer/discord-bot-go-template/internal/config"
)

type EventHandler struct {
	botClient     bot.Client
	messageRegex  *regexp.Regexp
	reactionEmoji string
}

func New(cfg *config.Config) *EventHandler {
	regex, err := regexp.Compile(cfg.MessageRegex)
	if err != nil {
		log.Fatalf("compiling regex: %v", err)
	}
	return &EventHandler{
		messageRegex:  regex,
		reactionEmoji: cfg.ReactionEmoji,
	}
}

func (h *EventHandler) InjectBotClient(client bot.Client) {
	h.botClient = client
}

func (h *EventHandler) OnReady(event *events.Ready) {
	log.Printf("connected as %s#%s\n", event.User.Username, event.User.Discriminator)
}

func (h *EventHandler) OnEvent(e bot.Event) {
	switch event := e.(type) {
	case *events.Ready:
		log.Printf("connected as %s#%s\n", event.User.Username, event.User.Discriminator)

	case *events.MessageCreate:
		fmt.Printf("MessageCreate triggered: %s\n", event.Message.Content)

		if event.Message.Author.ID == h.botClient.ApplicationID() {
			return
		}

		exampleReaction(h, event)

	case *events.ApplicationCommandInteractionCreate:
		switch event.Data.CommandName() {
		case "help":
			commands.HandleHelp(event, h.botClient)
		}

	}
}
