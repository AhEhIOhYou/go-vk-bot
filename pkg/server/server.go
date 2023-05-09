package server

import (
	"log"
	"math/rand"
	"time"

	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/constants"
	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/infrastructure"
	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/interfaces"
	"github.com/gin-gonic/gin"
)

func Serve() {

	rand.Seed(time.Now().Unix())

	// Creating services with repos
	services, err := infrastructure.NewRepo()
	if err != nil {
		log.Fatalf(constants.ServiceInitializationError, err)
	}

	// Creating event service
	event := interfaces.NewEventService(services.Vk, services.Nasa)

	router := gin.Default()

	// Routes processing
	router.POST("/", event.NewVkEvent)

	log.Fatal(router.Run())
}
