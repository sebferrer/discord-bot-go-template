package eventhandler

import (
	"log"

	"github.com/disgoorg/disgo/events"
)

func exampleReaction(h *EventHandler, event *events.MessageCreate) {
	if h.messageRegex.MatchString(event.Message.Content) {
		err := h.botClient.Rest().AddReaction(event.ChannelID, event.Message.ID, h.reactionEmoji)
		if err != nil {
			log.Printf("add reaction: %v", err)
		} else {
			log.Printf("reacted to the message '%s' with %s\n", event.Message.Content, h.reactionEmoji)
		}
	}
}
