package handlers

import (
	vrcbot "github.com/Jilwer/vrcgo/bot"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
)

func MessageHandler(b *vrcbot.Bot) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.MessageCreate) {
		// TODO: handle message
	})
}
