package main

import (
	"time"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func NewTempMessage(duration time.Duration, message discordgo.Message, afterChan chan discordgo.Message) *time.Timer {
	timer := time.AfterFunc(duration, func(){ afterChan <- message })
	return timer
}

func MessageDeleter(session *discordgo.Session, afterChan chan discordgo.Message) {
	for message := range afterChan {
		session.ChannelMessageDelete(message.ChannelID, message.ID)
	}
}

func getHighestRole(userRoles []string, guildRoles discordgo.Roles) *discordgo.Role {
	var highestRole *discordgo.Role = guildRoles[0]
	rolesMap := map[string]bool {}
	for _, id := range userRoles {
		rolesMap[id] = true
	}
	for _, role := range guildRoles {
		if rolesMap[role.ID] && role.Position > highestRole.Position {
			highestRole = role
		}
	}
	return highestRole
}

func getHighestRoleWithColor(userRoles []string, guildRoles discordgo.Roles) *discordgo.Role {
	var highestRole *discordgo.Role = guildRoles[0]
	rolesMap := map[string]bool {}
	for _, id := range userRoles {
		rolesMap[id] = true
	}
	for _, role := range guildRoles {
		if rolesMap[role.ID] && role.Position > highestRole.Position && role.Color != 0 {
			highestRole = role
		}
	}
	return highestRole
}

func getImagesFromAttachments(attachments []*discordgo.MessageAttachment) []*discordgo.MessageAttachment {
	result := []*discordgo.MessageAttachment {}
	for _, attachment := range attachments {
		if strings.HasPrefix(attachment.ContentType, "image") {
			result = append(result, attachment)
		}
	}
	return result
}

