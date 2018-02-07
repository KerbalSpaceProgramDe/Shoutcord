/*
 Shoutcord
 Bridging Modern Shoutbox and Discord

 Licensed under the Terms of the MIT License
 Copyright (c) 2018 KerbalSpaceProgram.de, Dorian Stoll
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/kataras/iris"
	"net/http"
	"github.com/spf13/cast"
)

// The connection to the discord server
var Discord *discordgo.Session

// The http server controlling the endpoint
var App *iris.Application

/*
 The entrypoint for the shoutcord application.
 Here we create the connection to discord and the endpoint for the shoutbox to submit new messages
 */
func main() {

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

	// Add the message handlers
	Discord.AddHandler(OnMessageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = Discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	defer Discord.Close()

	// Set up the endpoint
	App := iris.New()

	// Add the route
	App.Post("/message", OnMessageReceived)

	// Start listening
	App.Run(iris.Addr(Settings.LocalEndpoint))

}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages posted by bots, including ourselves
	if m.Author.Bot {
		return
	}

	// Ignore all messages outside of our specified channel
	if m.ChannelID != Settings.Channel {
		return
	}

	// I have no idea how exactly Text To Speech works in Discord, I think it is good to ignore it here
	if m.Tts {
		return
	}

	// Only transmit normal message, no fancy things
	if m.Type != discordgo.MessageTypeDefault {
		return
	}

	// Obtain the timestamp of the message
	t, err := m.Timestamp.Parse()
	if err != nil {
		panic(err)
	}

	// The message was valid, get the content and post it to the endpoint
	values := iris.Map {
		"username": m.Author.Username,
		"timestamp": t.Unix(),
		"message": m.ContentWithMentionsReplaced(),
	}
	jsonValue, _ := json.Marshal(values)
	_, err = http.Post(Settings.ShoutboxEndpoint, "application/json", bytes.NewBuffer(jsonValue))

	// If the endpoint had an error, quit
	if err != nil {
		panic(err)
	}

}

// This function will be called when a new message was created in the shoutbox
func OnMessageReceived(ctx iris.Context) {

	// Extract the JSON object from the request
	var values map[string]interface{}
	err := ctx.ReadJSON(&values)

	// Was there any error?
	if err != nil {
		panic(err)
	}

	// Fetch the data from the request
	// We don't need timestamp here, as we create our own one
	username, err := cast.ToStringE(values["username"])
	if err != nil {
		panic(err)
	}
	message, err := cast.ToStringE(values["message"])
	if err != nil {
		panic(err)
	}

	// Send a new message with the according data
	_, err = Discord.ChannelMessageSend(Settings.Channel, "`" + username + ":` " + message)
	if err != nil {
		panic(err)
	}

	// Answer the request
	ctx.WriteString("Ok")

}