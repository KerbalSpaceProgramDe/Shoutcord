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

/*
 Register the modules functions
*/
func init() {

    // Add the discord message handler
    app.Discord.AddHandler(onMessageCreate)

    // Reqister the job that polls new messages from the shoutbox
    gocron.Every(1).Second().Do(checkForNewMessages)
}
