/*
 Shoutcord
 Bridging Modern Shoutbox and Discord

 Licensed under the Terms of the MIT License
 Copyright (c) 2018 KerbalSpaceProgram.de, Dorian Stoll
*/

package main

import (
	"github.com/KerbalSpaceProgramDe/Shoutcord/app"

	// Import modules
	//_ "github.com/KerbalSpaceProgramDe/Shoutcord/modules/shoutbox"
	_ "github.com/KerbalSpaceProgramDe/Shoutcord/modules/welcome"
)

/*
 The entrypoint of our application. Instead of doing any significant work here,
 we delegate it to the app package
*/
func main() {
	app.Run()
}
