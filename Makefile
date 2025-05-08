# Detect OS and set delete command
ifeq ($(OS),Windows_NT)
    RM = cmd /c del
    BIN_EXT = .exe
else
    RM = rm -f
    BIN_EXT =
endif

# Name of the output binary
BINARY_NAME=discord-bot

# Path to the main application entry point
ENTRYPOINT=cmd/bot/main.go

.PHONY: all help build run clean deps start

# Default target
all: help

# Show available make targets
help:
	@echo "Available commands:"
	@echo "  make build    - Build the Go binary"
	@echo "  make start    - Build and run the saved binary"
	@echo "  make run      - Build and run the project without saving the binary"
	@echo "  make clean    - Remove the built binary"
	@echo "  make deps     - Ensure module dependencies are up to date"
	@echo "  make help     - Show this help message"

# Build the Go binary
build:
	go build -o bin/$(BINARY_NAME)$(BIN_EXT) $(ENTRYPOINT)

# Run the Go binary
start: build
	./bin/$(BINARY_NAME)$(BIN_EXT)

# Run the application
run:
	go run $(ENTRYPOINT)

# Clean up the binary
clean:
	$(RM) bin\\$(BINARY_NAME)$(BIN_EXT)

# Tidy Go modules
deps:
	go mod tidy
