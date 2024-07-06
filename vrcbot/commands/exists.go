package commands

import (
	"github.com/Jilwer/vrcgo/vrcapi"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var exists = discord.SlashCommandCreate{
	Name:        "exists",
	Description: "check if a email, displayname, or userid exists on the vrchat api",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionString{
			Name:         "type",
			Description:  "type of check to perform",
			Required:     true,
			Autocomplete: true,
		},
		discord.ApplicationCommandOptionString{
			Name:        "value",
			Description: "value to check",
		},
	},
}

func ExistsAutocompleteHandler(e *handler.AutocompleteEvent) error {
	return e.AutocompleteResult([]discord.AutocompleteChoice{
		discord.AutocompleteChoiceString{
			Name:  "Email",
			Value: "email",
		},
		discord.AutocompleteChoiceString{
			Name:  "Username",
			Value: "displayName",
		},
		discord.AutocompleteChoiceString{
			Name:  "UserID",
			Value: "userId",
		},
	})
}

func ExistsHandler(e *handler.CommandEvent) error {

	client, err := vrcapi.NewVRCApiClient(vrcapi.BaseURL, UserAgent)
	if err != nil {
		return err
	}

	filter := e.SlashCommandInteractionData().String("type")
	value := e.SlashCommandInteractionData().String("value")

	var exists bool
	switch filter {
	case vrcapi.FilterEmail:
		exists, err = client.CheckUserExists(filter, value)
	case vrcapi.FilterDisplayName:
		exists, err = client.CheckUserExists(filter, value)
	case vrcapi.FilterUserID:
		exists, err = client.CheckUserExists(filter, value)
	default:
		return e.CreateMessage(discord.NewMessageCreateBuilder().SetContent("Invalid filter type").Build())
	}
	if err != nil {
		return err
	}

	return e.CreateMessage(discord.NewMessageCreateBuilder().SetContentf("Exists: %t", exists).Build())
}
