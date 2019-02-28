/*
 Shoutcord
 Bridging Modern Shoutbox and Discord

 Licensed under the Terms of the MIT License
 Copyright (c) 2018 KerbalSpaceProgram.de, Dorian Stoll
*/

package emoji

import (
    "github.com/bwmarrin/discordgo"
    "regexp"
    "strings"
)

/*
 A regular expression used to find all discord formatted emojis.
 i.e. <:textcode:id>
*/
var emojiRegex = regexp.MustCompile(`<[a]?:([^:]+):\d+>`)

/*
 A regular expression to find all WBB formatted smileys.
 */
var forumRegex = regexp.MustCompile(`<img.*alt="([^"]+)".*>`)

/*
 Converts the unicode version of a discord emoji to the text code variant
*/
func escapeDiscordEmoji(input string) string {
    for emoji := range discordEmojiSwapped {
        input = strings.Replace(input, emoji, ":"+discordEmojiSwapped[emoji]+":", -1)
    }
    res := emojiRegex.FindAllStringSubmatch(input, -1)
    for _, item := range res {
        input = strings.Replace(input, item[0], ":"+item[1]+":", -1)
    }
    return input
}

/*
 Converts the discord emojis to forum text codes
*/
func DiscordToForumEmoji(input string) string {
    input = escapeDiscordEmoji(input)
    for emoji := range forumEmojiSwapped {
        input = strings.Replace(input, emoji, forumEmojiSwapped[emoji], -1)
    }
    return input
}

/*
 Converts the forum smiley codes to discord emojis
*/
func ForumToDiscordEmoji(input string, serverEmojis []*discordgo.Emoji) string {
    res := forumRegex.FindAllStringSubmatch(input, -1)
    for _,item := range res {
        input = strings.Replace(input, item[0], item[1], -1)
    }
    for emoji := range forumEmoji {
        for _, item := range serverEmojis {
            if forumEmoji[emoji] == ":"+item.Name+":" {
                input = strings.Replace(input, emoji, "<"+forumEmoji[emoji]+item.ID+">", -1)
            }
        }
    }
    return input
}
