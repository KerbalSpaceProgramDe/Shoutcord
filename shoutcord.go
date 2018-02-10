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
	"github.com/jasonlvhit/gocron"
	"github.com/spf13/cast"
	"net/http"
	"strings"
	"time"
)

// The connection to the discord server
var Discord *discordgo.Session

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

	// Setup the polling scheduler
	gocron.Every(1).Second().Do(CheckForNewMessages)
	<- gocron.Start()

}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages posted by ourselves
	if m.Author.ID == Discord.State.User.ID {
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

	// Only transmit normal messages, no fancy things
	if m.Type != discordgo.MessageTypeDefault {
		return
	}

	// Obtain the timestamp of the message
	t, err := m.Timestamp.Parse()
	if err != nil {
		panic(err)
	}

	// The message was valid, get the content and post it to the endpoint
	values := map[string]interface{} {
		"userID": 0,
		"username": m.Author.Username,
		"time": t.Unix(),
		"message": m.ContentWithMentionsReplaced(),
	}
	jsonValue, _ := json.Marshal(values)
	_, err = http.Post(Settings.Endpoint + "/shoutbox", "application/json", bytes.NewBuffer(jsonValue))

	// If the endpoint had an error, quit
	if err != nil {
		panic(err)
	}

}

// The current time. Will be updated every time the shoutbox is polled for new messages
var timestamp = time.Now().Unix()

// This function will check for new shoutbox messages every second
func CheckForNewMessages() {

	// Request data
	resp, err := http.Get(Settings.Endpoint + "/shoutbox?fromApi=0&since_time=" + cast.ToString(timestamp))

	// If the endpoint had an error, quit
	if err != nil {
		panic(err)
	}

	// Decode the request body
	var values []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&values)
	if err != nil {
		panic(err)
	}

	// Iterate over new messages
	for _,item := range values {

		// Fetch the data from the message
		// We don't need timestamp here, as we create our own one
		username, err := cast.ToStringE(item["username"])
		if err != nil {
			panic(err)
		}
		message, err := cast.ToStringE(item["message"])
		if err != nil {
			panic(err)
		}

		// Remove BBCode from the message
		message = strings.Replace(message, "[url]", "", 0)
		message = strings.Replace(message, "[/url]", "", 0)

		// Send a new message with the according data
		_, err = Discord.ChannelMessageSend(Settings.Channel, "`" + username + ":` " + message)
		if err != nil {
			panic(err)
		}
	}

	// Update the timestamp
	timestamp = time.Now().Unix()

}