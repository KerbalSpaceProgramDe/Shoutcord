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
	// Request data
	resp, err := utils.GetHTTP(app.Settings.Endpoint+"/shoutbox?fromApi=0&since_time="+cast.ToString(shoutboxTimestamp),
		app.Settings.ApiKey)

	// If the endpoint had an error, quit
	if err != nil {
		panic(err)
	}

	// Decode the request body
	var temp struct {
		Data []map[string]interface{} `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&temp)
	if err != nil {
		panic(err)
	}
	values := temp.Data

	// Iterate over new messages
	for _, item := range values {

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

		// Remove BBCodes from the message
		message = utils.RemoveLinkCodes(message)

		// Replace the forum smileys with their Kerbal Emoji version
		message = emoji.ForumToDiscordEmoji(message, app.Guild.Emojis)

		// Send a new message with the according data
		_, err = app.Discord.ChannelMessageSend(app.Settings.Channel, "`"+username+":` "+message)
		if err != nil {
			panic(err)
		}

		// Update the timestamp
		if mTimestamp > shoutboxTimestamp {
			shoutboxTimestamp = mTimestamp
		}
	}
}
