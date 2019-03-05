/*
 Shoutcord
 Bridging Modern Shoutbox and Discord

 Licensed under the Terms of the MIT License
 Copyright (c) 2018 KerbalSpaceProgram.de, Dorian Stoll
*/

package shoutbox

import (
	"bytes"
	"encoding/json"
	"github.com/KerbalSpaceProgramDe/Shoutcord/app"
	"github.com/KerbalSpaceProgramDe/Shoutcord/emoji"
	"github.com/KerbalSpaceProgramDe/Shoutcord/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/cast"
)

/*
 This function will be called every time a new message is created
 on any channel that the autenticated bot has access to.
*/
func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages posted by ourselves
	if m.Author.ID == app.Discord.State.User.ID {
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

	// Get a list of all attachments and transform them to their URLs
	attachments := ""
	for _, item := range m.Attachments {
		attachments = attachments + item.URL + " "
	}

	// Replace the Kerbal Emojis with their forum version
	message := emoji.DiscordToForumEmoji(m.ContentWithMentionsReplaced())

	// Format links so they get displayed correctly in the shoutbox
	message = utils.AddLinkCodes(message + " " + attachments)

	// Obtain the timestamp of the message
	t, err := m.Timestamp.Parse()
	if utils.HandleError(err) {
		return
	}

	// The message was valid, get the content and post it to the endpoint
	jsonValue, _ := json.Marshal(map[string]interface{}{
		"discordId": cast.ToInt(m.Author.ID),
		"username":  m.Author.Username,
		"time":      t.Unix(),
		"message":   message,
	})
	_, err = utils.PostHTTP(Settings.Endpoint+"/shoutbox", "application/json", bytes.NewBuffer(jsonValue),
		Settings.ApiKey)

	// If the endpoint had an error, quit
	if utils.HandleError(err) {
		return
	}
}
