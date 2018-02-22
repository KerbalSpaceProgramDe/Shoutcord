/*
 Shoutcord
 Bridging Modern Shoutbox and Discord

 Licensed under the Terms of the MIT License
 Copyright (c) 2018 KerbalSpaceProgram.de, Dorian Stoll
*/

package app

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

/*
 Starts a simple webserver that responds with pong to everything
*/
func ServePingPong(host string) error {
	http.HandleFunc("/", handler)
	return http.ListenAndServe(host, nil)
}
