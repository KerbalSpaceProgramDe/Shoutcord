/*
 Shoutcord
 Bridging Modern Shoutbox and Discord

 Licensed under the Terms of the MIT License
 Copyright (c) 2018 KerbalSpaceProgram.de, Dorian Stoll
*/

package utils

import (
	"github.com/KerbalSpaceProgramDe/Shoutcord/app"
	"log"
)

/*
 Handles an error message by terminating the program in debug mode
*/
func HandleError(err error) bool {
	if app.Settings.Debug && err != nil {
		panic(err)
	}
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
