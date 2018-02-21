/*
 Shoutcord
 Bridging Modern Shoutbox and Discord

 Licensed under the Terms of the MIT License
 Copyright (c) 2018 KerbalSpaceProgram.de, Dorian Stoll
*/

package shoutbox

import (
	"encoding/json"
	"github.com/KerbalSpaceProgramDe/Shoutcord/app"
	"github.com/KerbalSpaceProgramDe/Shoutcord/emoji"
	"github.com/KerbalSpaceProgramDe/Shoutcord/utils"
	"github.com/spf13/cast"
	"time"
)

/*
 The current time. Will be updated every time the shoutbox is polled for new messages
*/
var shoutboxTimestamp = time.Now().Unix()

/*
 This function will check for new shoutbox messages every second
*/
func checkForNewMessages() {

	// Try to figure out the server we are in
	if Guild == nil {

		// Fetch the channel we are in
		channel, err := app.Discord.Channel(Settings.Channel)
		if utils.HandleError(err) {
			return
		}

		// Fetch the server of the channel
		guild, err := app.Discord.Guild(channel.GuildID)
		if utils.HandleError(err) {
			return
		}
		Guild = guild
	}

	// Request data
	resp, err := utils.GetHTTP(Settings.Endpoint+"/shoutbox?fromApi=0&since_time="+cast.ToString(shoutboxTimestamp),
		Settings.ApiKey)

	// If the endpoint had an error, quit
	if utils.HandleError(err) {
		return
	}

	// Decode the request body
	var temp struct {
		Data []map[string]interface{} `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&temp)
	if utils.HandleError(err) {
		return
	}
	values := temp.Data

	// Iterate over new messages
	for _, item := range values {

		// Fetch the data from the message
		username, err := cast.ToStringE(item["username"])
		if utils.HandleError(err) {
			return
		}
		message, err := cast.ToStringE(item["message"])
		if utils.HandleError(err) {
			return
		}
		mTimestamp, err := cast.ToInt64E(item["time"])
		if utils.HandleError(err) {
			return
		}

		// Remove BBCodes from the message
		message = utils.RemoveLinkCodes(message)

		// Replace the forum smileys with their Kerbal Emoji version
		message = emoji.ForumToDiscordEmoji(message, Guild.Emojis)

		// Send a new message with the according data
		_, err = app.Discord.ChannelMessageSend(Settings.Channel, "`"+username+":` "+message)
		if utils.HandleError(err) {
			return
		}

		// Update the timestamp
		if mTimestamp > shoutboxTimestamp {
			shoutboxTimestamp = mTimestamp
		}
	}
}
