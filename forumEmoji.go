/*
 Shoutcord
 Bridging Modern Shoutbox and Discord

 Licensed under the Terms of the MIT License
 Copyright (c) 2018 KerbalSpaceProgram.de, Dorian Stoll
 */

 package main

import (
    "strings"
)

// A list of all forum emojis and their corresponding discord emojis
var forumEmoji = map[string]string {
    ":)": ":kerbsmile:",
    ":P": ":kerbtongue:",
    ":(": ":kerbsad:",
    ";)": ":kerbwink:",
    "8)": ":kerbcool:",
    ":D": ":kerbbiggrin:",
    ";(": ":kerbcrylaugh:",
    "^^": ":kerbsquint:",
    ":rolleyes:": ":kerbrolleyes:",
    ":huh:": ":kerbhuh:",
    ":S": ":kerbunsure:",
    ":love:": ":kerblove:",
    ":seriously:": ":kerbseriously:",
    ":thumbup:": ":kerbthumbup:",
    ":thumbsup:": ":kerbthumbsup:",
    ":thumbdown:": ":kerbthumbdown:",
    ":|": ":kerbmellow:",
    "8o": ":kerbw00t:",
    ":pinch:": ":kerbpinch:",
    ":sleeping:": ":kerbsleeping:",
    ":wacko:": ":kerbwacko:",
    ";P": ":kerbwinktongue:",
    ":]": ":kerbembarressed:",
    ":party:": ":kerbparty:",
    ":whistling:": ":kerbwhistling:",
    ":cursing:": ":kerbcursing:",
    ":crylaugh:": ":kerbcrylaugh:",
    ":facepalm:": ":kerbfacepalm:",
    ":really:": ":kerbreally:",
    ":badplan:": ":kerbbadplan:",
    ":wink:": ":kerbwave:",
    ":evil:": ":kerbevil:",
    ":saint:": ":kerbsaint:",
    ":kerbonaut:": ":kerbonaut:",
    ":ahhh:": ":kerbthinking",
    ":funds:": ":kerbfunds:",
    ":Funds1:": ":kerbfunds1:",
    ":science:": ":kerbscience:",
    ":science1:": ":kerbscience1:",
    ":Pudding:": ":kerbpudding:",
    ":nerd:": ":kerbnerd:",
    "xP": ":kerbpinchtongue:",
    ":finger:": ":kerbfinger:",
    ":krank:": ":kerbsick:",
    ":ekeks:": ":kerbcookie:",
    ":schäm:": ":kerbshaming:",
    ":gebdat:": ":kerbbirthday:",
    ":Pilot:": ":kerbpilot:",
    "-.-": ":kerbgrumply:",
    ":OsterKerbal:": ":kerbrabbit:",
    ":franky:": ":franky:",
    ":Hell-o-Week:": ":kerbpumpkin:",
    ":wmann:": ":kerbwmann:",
    ":wbaum:": ":kerbwbaum:",
    ":lichtrot:": ":kerblightred:",
    ":lichtgelb:": ":kerblightyellow:",
    ":lichtgrün:": ":kerblightgreen:",
    ":lichtblau:": ":kerblightblue:",
}
var forumEmojiSwapped = SwapMap(forumEmoji)

func DiscordToForumEmoji(message string) string {
    for emoji := range forumEmojiSwapped {
        message = strings.Replace(message, emoji, forumEmojiSwapped[emoji], -1)
    }
    return message
}

func ForumToDiscordEmoji(message string) string {
    for emoji := range forumEmoji {
        for _,item := range Guild.Emojis {
            if forumEmoji[emoji] == ":" + item.Name + ":" {
                message  = strings.Replace(message, emoji, "<" + forumEmoji[emoji] + item.ID + ">", -1)
            }
        }
    }
    return message
}
