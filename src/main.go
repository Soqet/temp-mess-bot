package main

import (
	"fmt"
	"os"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"tmp-mess-bot/pkg/logger"
)

var (
	AfterChan chan discordgo.Message
	//LogChan chan<- string
	LogString func(string)
)

func initEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Print("No .env file found")
	}
}

func main() {
	initEnv()
	stopChan := make(chan any)
	var botToken = os.Getenv("TOKEN")
	discord, err := discordgo.New("Bot " + botToken)
	if err != nil { return }

	logger := logger.Logger{}
	logFile, err := os.OpenFile("logs.txt", os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil { fmt.Println("huy"); return }
	defer logFile.Close()
	logChan := make(chan []byte)
	logger.Init(logFile, logChan)
	LogChan := byteChannelAdapter(logChan)
	LogString = func(s string) {
		fmt.Println(s)
		LogChan <- s
	}
	AfterChan = make(chan discordgo.Message)
	go MessageDeleter(discord, AfterChan)

	discord.AddHandler(messageCreate)
	discord.Open()
	defer discord.Close()
	go func (schan chan any) {
		consoleCommand := ""
		for consoleCommand != "stop" {
			fmt.Scanln(&consoleCommand)
		}
		close(schan)
	}(stopChan)
	<-stopChan
}