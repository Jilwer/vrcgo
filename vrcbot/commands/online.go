package commands

import (
	"github.com/Jilwer/vrcgo/vrcapi"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var online = discord.SlashCommandCreate{
	Name:        "online",
	Description: "return the number of users online in vrchat",
}

func OnlineHandler(e *handler.CommandEvent) error {

	vrcapiClient, err := vrcapi.NewVRCApiClient(vrcapi.BaseURL, "vrcgo/0.1 rustup")
	if err != nil {
		return err
	}

	online, err := vrcapiClient.GetOnlineUsers()
	if err != nil {
		return err
	}

	return e.CreateMessage(discord.NewMessageCreateBuilder().
		SetContentf("Online: %s", online).Build(),
	)
}
