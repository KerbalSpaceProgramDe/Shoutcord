/*
 Shoutcord
 Bridging Modern Shoutbox and Discord

 Licensed under the Terms of the MIT License
 Copyright (c) 2018 KerbalSpaceProgram.de, Dorian Stoll
*/

package utils

import (
    "io"
    "net/http"
)

func GetHTTP(url string, key string) (resp *http.Response, err error) {
    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Set("X-Authorization", key)
    return client.Do(req)
}

func PostHTTP(url string, bodyType string, data io.Reader, key string) (resp *http.Response, err error) {
    client := &http.Client{}
    req, err := http.NewRequest("POST", url, data)
    req.Header.Set("Content-Type", bodyType)
    req.Header.Set("X-Authorization", key)
    return client.Do(req)
}
