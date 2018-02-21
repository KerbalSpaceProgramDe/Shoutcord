/*
 Shoutcord
 Bridging Modern Shoutbox and Discord

 Licensed under the Terms of the MIT License
 Copyright (c) 2018 KerbalSpaceProgram.de, Dorian Stoll
*/

package app

import (
	"github.com/jinzhu/configor"
	"log"
	"os"
)

/*
 All the config variables from the config file
*/
var Settings = struct {

	// The token for the bot user
	Token string `yaml:"token"`

	// Whether the bot should run in debug mode
	Debug bool `yaml:"debug"`
}{}

/*
 Loads the settings from the config file
*/
func LoadSettings() {
	LoadFromConfigFile(&Settings, "config.yml")
}

/*
 Loads the settings from a configuration file
*/
func LoadFromConfigFile(data interface{}, configFile string) {
	log.Printf("* Loading configuration file: config/%s", configFile)
	os.Setenv("CONFIGOR_ENV_PREFIX", "SHOUTCORD")
	err := configor.Load(data, "config/"+configFile)
	if err != nil {
		log.Fatalf("* Failed to parse configuration file: %s", err)
	}
}
