package server

import (
	"log"

	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/interfaces"
	"github.com/gin-gonic/gin"
)

func Serve() {
	event := interfaces.NewEvent()

	router := gin.Default()

	router.POST("/", event.Index)

	log.Fatal(router.Run())
}
