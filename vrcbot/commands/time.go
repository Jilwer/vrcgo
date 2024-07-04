package commands

import (
	"github.com/Jilwer/vrcgo/vrcapi"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var time = discord.SlashCommandCreate{
	Name:        "time",
	Description: "return the system time of vrchat",
}

func TimeHandler(e *handler.CommandEvent) error {

	vrcapiClient, err := vrcapi.NewVRCApiClient(vrcapi.BaseURL, UserAgent)
	if err != nil {
		return err
	}

	time, err := vrcapiClient.GetSystemTime()
	if err != nil {
		return err
	}

	return e.CreateMessage(discord.NewMessageCreateBuilder().
		SetContentf("System Time: %s", time).Build(),
	)
}
