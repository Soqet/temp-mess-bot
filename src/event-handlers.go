package main

import ( 
	"strings"
	"github.com/bwmarrin/discordgo"
)


func messageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.Bot || !strings.HasPrefix(event.Content, commandPrefix) { return }
	firstSpace := strings.IndexByte(event.Content, ' ')
	var command string
	if firstSpace == -1 {
		command = strings.Clone(event.Content[len(commandPrefix):])
	} else {
		command = strings.Clone(event.Content[len(commandPrefix):firstSpace])
	}
	switch command {
	case commTmp:
		go tmpMess(session, event)
	default:
		unknownCommand(session, event)
	}
	
}