package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

type Command struct {
	Definition          discord.ApplicationCommandCreate
	Handler             handler.CommandHandler
	AutoCompleteHandler handler.AutocompleteHandler
}

var Commands = []Command{
	{
		test,
		TestHandler,
		TestAutocompleteHandler,
	},
	{
		online,
		OnlineHandler,
		nil,
	},
	{
		config,
		ConfigHandler,
		nil,
	},
	{
		time,
		TimeHandler,
		nil,
	},
	{
		exists,
		ExistsHandler,
		ExistsAutocompleteHandler,
	},
}
