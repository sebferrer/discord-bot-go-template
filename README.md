# Discord Bot Go Template

A minimal and modular architecture for building Discord bots in Go using the [Disgo](https://github.com/disgoorg/disgo) library.

## Features

- Support for slash commands (e.g., `/help`)
- Event handling for messages and reactions
- Environment-based configuration using `.env` files
- Makefile for streamlined build and run processes
- Support for multiple environments (`dev` and `prod`)

## Project Structure

```
discord-bot-go-template/
├── .env                # Environment variables
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
├── README.md
├── cmd/
│   └── bot/
│       └── main.go     # Entry point of the application
└── internal/
    ├── bot/
    │   ├── bot.go      # Bot client setup and connection
    │   ├── commands/
    │   │   ├── help.go     # /help command implementation
    │   │   └── manager.go  # Command registration logic
    │   └── eventhandler/
    │       ├── event-handler.go    # Event handling logic
    │       └── example-reaction.go # Example reaction handler
    └── config/
        └── config.go   # Configuration loading
```

## Getting Started

### Prerequisites

- Go 1.23 or higher
- A Discord application with a bot token
- A Discord server where you have permission to add bots

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/sebferrer/discord-bot-go-template.git
   cd discord-bot-go-template
   ```

2. **Set up environment variables:**

   Create a `.env` file in the root directory with the following content:

   ```env
   DISCORD_BOT_TOKEN=your_bot_token_here
   DISCORD_CLIENT_ID=your_client_id_here
   DISCORD_GUILD_ID=your_guild_id_here
   ENV=dev
   ```

   - `DISCORD_BOT_TOKEN`: Your bot's token from the Discord Developer Portal.
   - `DISCORD_CLIENT_ID`: Your application's unique identifier on Discord, used for bot authorization and OAuth2 flows. 
   - `DISCORD_GUILD_ID`: The ID of your Discord server (used in development mode).
   - `ENV`: Set to `dev` for development mode or `prod` for production mode.

3. **Run the bot:**

- Run directly without saving the binary (temporary build):

  ```bash
  make run
  ```

- Build a binary and run it:

  ```bash
  make start
  ```

> `make start` automatically triggers `make build` under the hood.

You can also run the build step manually:

```bash
make build
```

This will compile the bot into `bin/discord-bot(.exe)`, which you can run directly if needed.

## Usage

Once the bot is running and added to your Discord server, you can use the following command:

- `/help`: Lists all available commands.

## Development

### Adding New Commands

1. Create a new file in `internal/bot/commands/`, e.g., `ping.go`.
2. Implement the command logic following the structure in `help.go`.
3. Register the new command in `manager.go`.

### Event Handling

To handle new events:

1. Add a new handler function in `internal/bot/eventhandler/`.
2. Register the handler in `event-handler.go`.

## Makefile Commands

```makefile
make help    # Display available make commands
make build   # Build the Go binary to bin/discord-bot
make start   # Build and run the saved binary from ./bin
make run     # Build and run the project without saving the binary
make clean   # Remove the built binary
make deps    # Ensure module dependencies are up to date
```

## License

This project is licensed under the [Apache-2.0 License](LICENSE).
