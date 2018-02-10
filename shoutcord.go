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
	"mvdan.cc/xurls"
	"net/http"
	"regexp"
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

	// Get a list of all attachments
	attachments := ""
	for _,item := range m.Attachments {
		attachments = attachments + "[url]" + item.URL + "[/url] "
	}

	// Detect URLs
	message := m.ContentWithMentionsReplaced()
	matches := removeDuplicatesUnordered(xurls.Strict.FindAllString(message, -1))
	for _,item := range matches {
		message = strings.Replace(message, item, "[url]" + item + "[/url]", -1)
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
		"message": message + " " + attachments,
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

// The regex that is used to detect links in the shoutbox messages
var re = regexp.MustCompile(`\[url(?:=(?:["']?)([^"'\]]+)(?:["']?))?]([^\[]+)\[\/url\]`)

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
		username, err := cast.ToStringE(item["username"])
		if err != nil {
			panic(err)
		}
		message, err := cast.ToStringE(item["message"])
		if err != nil {
			panic(err)
		}
		mTimestamp, err := cast.ToInt64E(item["time"])
		if err != nil {
			panic(err)
		}

		// Remove BBCode from the message
		res := re.FindAllStringSubmatch(message, -1)
		for _,item := range res {
			link := item[1]
			desc := item[2]

			// If the link is empty, it was stored in the description value
			if link == "" {
				link = desc
				desc = ""
			}

			// If a description exists, this is a relative link
			if strings.HasPrefix(desc, "@") {
				message = strings.Replace(message, item[0], desc, -1)
			} else {
				message = strings.Replace(message, item[0], link, -1)
			}
		}

		// Send a new message with the according data
		_, err = Discord.ChannelMessageSend(Settings.Channel, "`" + username + ":` " + message)
		if err != nil {
			panic(err)
		}

		// Update the timestamp
		if mTimestamp > timestamp {
			timestamp = mTimestamp
		}
	}

}

// Taken from: https://www.dotnetperls.com/duplicates-go
func removeDuplicatesUnordered(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v:= range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}