/*
 Shoutcord
 Bridging Modern Shoutbox and Discord

 Licensed under the Terms of the MIT License
 Copyright (c) 2018 KerbalSpaceProgram.de, Dorian Stoll
*/

package app

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jasonlvhit/gocron"
)

/*
 The connection to the discord API
*/
var Discord *discordgo.Session

/*
 The discord server we are active in
*/
var Guild *discordgo.Guild

/*
 Load the config files and prepare the connection to discord
*/
func init() {
	// Load the settings file
	LoadSettings()

	// Connect to discord
	session, err := discordgo.New("Bot " + Settings.Token)

	// Any error?
	if err != nil {
		panic(err)
	}

	// Assign the session
	Discord = session
}

/*
 Open the connection to discord and start the job scheduler
*/
func Run() {
	// Open a websocket connection to Discord and begin listening.
	err := Discord.Open()
	if err != nil {
		panic(err)
	}
	defer Discord.Close()

	// Fetch the channel we are in
	channel, err := Discord.Channel(Settings.Channel)
	if err != nil {
		panic(err)
	}

	// Fetch the server of the channel
	guild, err := Discord.Guild(channel.GuildID)
	if err != nil {
		panic(err)
	}
	Guild = guild

	// Start the job scheduler
	<-gocron.Start()
}
