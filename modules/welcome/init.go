/*
 Shoutcord
 Bridging Modern Shoutbox and Discord

 Licensed under the Terms of the MIT License
 Copyright (c) 2018 KerbalSpaceProgram.de, Dorian Stoll
*/

package welcome

import (
	"github.com/KerbalSpaceProgramDe/Shoutcord/app"
	"github.com/KerbalSpaceProgramDe/Shoutcord/utils"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var Settings = struct {

	// The channel the welcome / goodbye announcements should get posted in.
	Channel string `yaml:"channel"`

	// The message that is sent when the user joins
	WelcomeMessage string `yaml:"welcome-message"`

	// The message that is sent when the user leaves
	GoodbyeMessage string `yaml:"goodbye-message"`

	// Additional message components that should get added after the welcome message
	WelcomeHints []string `yaml:"welcome-hints"`

	// Additional message components that should get added after the welcome message
	GoodbyeHints []string `yaml:"goodbye-hints"`
}{}

/*
 Register the modules functions
*/
func init() {

	// Load the settings
	app.LoadFromConfigFile(&Settings, "welcome.yml")

	// Add the discord message handler
	app.Discord.AddHandler(onGuildMemberAdded)
	app.Discord.AddHandler(onGuildMemberRemoved)
}

/*
 Gets called when a user joins the server
*/
func onGuildMemberAdded(session *discordgo.Session, m *discordgo.GuildMemberAdd) {

	// Get the welcome message and format it
	message := strings.Replace(Settings.WelcomeMessage, "{USERNAME}", m.User.Mention(), -1)

	// Append a random hint
	message += " " + Settings.WelcomeHints[utils.Random(0, len(Settings.WelcomeHints))]

	// Send the message
	_, err := app.Discord.ChannelMessageSend(Settings.Channel, message)
	if err != nil {
		panic(err)
	}
}

/*
 Gets called when a user leaves the server
*/
func onGuildMemberRemoved(session *discordgo.Session, m *discordgo.GuildMemberRemove) {

	// Get the goodbye message and format it
	message := strings.Replace(Settings.GoodbyeMessage, "{USERNAME}", m.User.Username, -1)

	// Append a random hint
	message += " " + Settings.GoodbyeHints[utils.Random(0, len(Settings.GoodbyeHints))]

	// Send the message
	_, err := app.Discord.ChannelMessageSend(Settings.Channel, message)
	if err != nil {
		panic(err)
	}
}
