/*
 Shoutcord
 Bridging Modern Shoutbox and Discord

 Licensed under the Terms of the MIT License
 Copyright (c) 2018 KerbalSpaceProgram.de, Dorian Stoll
*/

package shoutbox

import (
	"github.com/KerbalSpaceProgramDe/Shoutcord/app"
	"github.com/jasonlvhit/gocron"
)

var Settings = struct {

	// The channel that should get synced with the shoutbox
	Channel string `yaml:"channel"`

	// The API Endpoint that should receive the updates from discord
	Endpoint string `yaml:"endpoint"`

	// The access token for the kerbal.de API
	ApiKey string `yaml:"api-key"`
}{}

/*
 Register the modules functions
*/
func init() {

	// Load the settings
	app.LoadFromConfigFile(&Settings, "shoutbox.yml")

	// Add the discord message handler
	app.Discord.AddHandler(onMessageCreate)

	// Reqister the job that polls new messages from the shoutbox
	gocron.Every(1).Second().Do(checkForNewMessages)
}
