package main

import (
	"fmt"
	"time"
)

func byteChannelAdapter(bchan chan<- []byte) chan<- string {
	schan := make(chan string)
	go func(){
		for mess := range schan {
			bchan <- []byte(mess)
		}
		close(bchan)
	}()
	return schan
}


func formatLog(action string, author string, where string, result string) string {
	return fmt.Sprintf("%s\t%s by %s\tin %s\t%s\n", time.Now(), action, author, where, result)
}

