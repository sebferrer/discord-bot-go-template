package commands

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func HandleHelp(event *events.ApplicationCommandInteractionCreate, client bot.Client) {
	event.CreateMessage(discord.MessageCreate{
		Content: "Available commands:\n- /help: List all available commands",
	})
}
