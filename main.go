package main

import (
	"log"

	"github.com/AhEhIOhYou/go-vk-bot/pkg/server"
)

func main() {
	log.Println("hello vk api")

	server.Serve()
}
