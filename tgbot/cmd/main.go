package main

import (
	"log"
	"os"
	"tgbot/internal/core"
)

func main() {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		panic("BOT_TOKEN variable must be set")
	}

	endpoint := "http://comix-app-service:8080/api/v1/search"
	log.Println("Server is running")
	core.Run(botToken, endpoint)
}
