package main

import (
	"github.com/bwmarrin/discordgo"
	"tmp-mess-bot/pkg/parse"
	"time"
	"fmt"
)

func unknownCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	LogString(formatLog("Unknown command", message.Author.ID, message.GuildID + ":" + message.ID, ""))
	session.ChannelMessageSend(message.ChannelID, "Unknown command.")
}

func tmpMess(session *discordgo.Session, message *discordgo.MessageCreate) {
	var result string
	defer func(){ LogString(formatLog("tmpMess command", message.Author.ID, message.GuildID + ":" + message.ID, result)) }()
	duration, err := time.ParseDuration(parse.ParseWord(message.Content, 1))
	if err != nil { 
		result = "parse error"
		return 
	}
	session.ChannelMessageDelete(message.ChannelID, message.ID)
	guild, err := session.Guild(message.GuildID)
	if err != nil { 
		result = "guild error"
		return 
	}
	highestRoleWithColor := getHighestRoleWithColor(message.Member.Roles, guild.Roles)
	images := getImagesFromAttachments(message.Attachments)
	embed := &discordgo.MessageEmbed{
    Author:      &discordgo.MessageEmbedAuthor{},
    Color:       highestRoleWithColor.Color,
    Description: parse.DeleteFirstNWords(message.Content, 2),
    Thumbnail: &discordgo.MessageEmbedThumbnail{
        URL: message.Author.AvatarURL("2048"),
    },
    Timestamp: time.Now().Add(duration).Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
    Title:     fmt.Sprintf("%s#%s", message.Author.Username, message.Author.Discriminator),
	}
	if len(images) > 0 {
		embed.Image = &discordgo.MessageEmbedImage{
			URL: images[0].URL,
		}
	}
	botMess, err := session.ChannelMessageSendEmbed(message.ChannelID, embed)
	if err != nil { 
		result = "cant send message"
		return
 	}
	result = "success" 
	NewTempMessage(duration, *botMess, AfterChan)
}
