/*
 Shoutcord
 Bridging Modern Shoutbox and Discord

 Licensed under the Terms of the MIT License
 Copyright (c) 2018 KerbalSpaceProgram.de, Dorian Stoll
*/

package utils

import (
    "mvdan.cc/xurls"
    "regexp"
    "strings"
)

/*
 Converts discord links into a BBCode representation that works in the shoutbox
*/
func AddLinkCodes(input string) string {
    matches := RemoveDuplicatesUnordered(xurls.Strict.FindAllString(input, -1))
    for _, item := range matches {
        input = strings.Replace(input, item, "[url]"+item+"[/url]", -1)
    }
    return input
}

/*
 The regex that is used to detect links in the shoutbox messages
*/
var links = regexp.MustCompile(`\[url(?:=(?:["']?)([^"'\]]+)(?:["']?))?]([^\[]+)\[\/url\]`)

/*
 Removes the BBCode links from shoutbox posts so the link render correctly in discord
*/
func RemoveLinkCodes(input string) string {
    res := links.FindAllStringSubmatch(input, -1)
    for _, item := range res {
        link := item[1]
        desc := item[2]

        // If the link is empty, it was stored in the description value
        if link == "" {
            link = desc
            desc = ""
        }

        // If a description exists, this is a relative link
        if strings.HasPrefix(desc, "@") {
            input = strings.Replace(input, item[0], desc, -1)
        } else {
            input = strings.Replace(input, item[0], link, -1)
        }
    }
    return input
}
