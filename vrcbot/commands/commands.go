package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

type Command struct {
	Definition discord.ApplicationCommandCreate
	Handler    handler.CommandHandler
}

var Commands = []Command{
	{
		test,
		TestHandler,
	},
	{
		online,
		OnlineHandler,
	},
	{
		config,
		ConfigHandler,
	},
	{
		time,
		TimeHandler,
	},
}

//var Commands = []discord.ApplicationCommandCreate{
//	test,
//	version,
//	online,
//	config,
//}
