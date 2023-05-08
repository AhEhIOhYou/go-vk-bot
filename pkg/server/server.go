package server

import (
	"fmt"
	"log"

	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/constants"
	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/infrastructure"
	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/interfaces"
	"github.com/gin-gonic/gin"
)

func Serve() {

	services, err := infrastructure.NewRepo()
	if err != nil {
		log.Fatal(fmt.Sprintf(constants.ServiceInitializationError, err))
	}

	event := interfaces.NewEventService(services.Vk, services.Nasa)

	router := gin.Default()

	router.POST("/", event.NewMessage)

	log.Fatal(router.Run())
}
