package commands

import (
	"log"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
)

var commandList = []discord.ApplicationCommandCreate{
	discord.SlashCommandCreate{
		Name:        "help",
		Description: "List all available commands",
	},
}

func RegisterCommandsForGuild(client bot.Client, guildID snowflake.ID) {
	_, err := client.Rest().SetGuildCommands(
		client.ApplicationID(),
		guildID,
		commandList,
	)
	if err != nil {
		log.Fatalf("register guild commands: %v", err)
	}
}

func RegisterGlobalCommands(client bot.Client) {
	_, err := client.Rest().SetGlobalCommands(
		client.ApplicationID(),
		commandList,
	)
	if err != nil {
		log.Fatalf("register global commands: %v", err)
	}
}
