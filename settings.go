/*
 Shoutcord
 Bridging Modern Shoutbox and Discord

 Licensed under the Terms of the MIT License
 Copyright (c) 2018 KerbalSpaceProgram.de, Dorian Stoll
 */

package main

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

	// The channel that should get synced with the shoutbox
	Channel string `yaml:"channel"`

	// The API Endpoint that should receive the updates from discord
	Endpoint string `yaml:"endpoint"`

	// The access token for the kerbal.de API
	ApiKey string `yaml:"api-key"`

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
	log.Printf("* Loading configuration file: %s", configFile)
	os.Setenv("CONFIGOR_ENV_PREFIX", "SHOUTCORD")
	err := configor.Load(data, configFile)
	if err != nil {
		log.Fatalf("* Failed to parse configuration file: %s", err)
	}
}