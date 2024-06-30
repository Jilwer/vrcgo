package commands

import (
	"encoding/json"
	"strings"

	"github.com/Jilwer/vrcgo/vrcapi"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var config = discord.SlashCommandCreate{
	Name:        "config",
	Description: "return the system config for VRChat",
}

func ConfigHandler(e *handler.CommandEvent) error {

	vrclient, err := vrcapi.NewVRCApiClient(vrcapi.BaseURL, UserAgent)
	if err != nil {
		return err
	}

	config, err := vrclient.GetSystemConfig()
	if err != nil {
		return err
	}

	// marshal the config to json
	configJson, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return e.CreateMessage(discord.NewMessageCreateBuilder().
		SetContent("Here is the system config for VRChat").AddFile("config.json", "vrchat system config", strings.NewReader(string(configJson))).Build(),
	)
}
