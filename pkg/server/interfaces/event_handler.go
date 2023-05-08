package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/constants"
	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/entities"
	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/infrastructure"
	"github.com/gin-gonic/gin"
)

type ConfirmPost struct {
	Type    string `json:"type"`
	GroupID int    `json:"group_id"`
}

type EventService struct {
	VkApp   infrastructure.VkRepo
	NasaApp infrastructure.NasaRepo
}

func NewEventService(vkRepo infrastructure.VkRepo, nasaRepo infrastructure.NasaRepo) *EventService {
	return &EventService{
		VkApp:   vkRepo,
		NasaApp: nasaRepo,
	}
}

func (e *EventService) NewMessage(c *gin.Context) {

	var data *entities.Event

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf(constants.RequestFailed, err))
		log.Println(err)
	}

	// Object can be different types, so i use interface

	switch data.Type {
	case "message_new":
		var messageNew entities.MessageNew
		jsonData, _ := json.Marshal(data.Object)
		json.Unmarshal(jsonData, &messageNew)
		fmt.Println("Message text:")
		fmt.Println(messageNew.Message.Text)

		err := e.VkApp.SendMessage(&entities.MessageResponse{
			Message: "test-response-8",
			UserID:  messageNew.Message.PeerID,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf(constants.RequestFailed, err))
			log.Println(err)
		}
	case "message_event":
		var messageEvent entities.MessageEvent
		jsonData, _ := json.Marshal(data.Object)
		json.Unmarshal(jsonData, &messageEvent)
		fmt.Println("Event id:")
		fmt.Println(messageEvent.EventID)
		fmt.Println("Event payload:")
		fmt.Println(messageEvent.Payload)

		err := e.VkApp.SendEvent(&entities.EventResponse{
			EventID: messageEvent.EventID,
			UserID:  messageEvent.UserID,
			PeerID:  messageEvent.PeerID,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf(constants.RequestFailed, err))
			log.Println(err)
		}
	}

	c.Data(http.StatusOK, "charset=utf8", []byte("ok"))
}

func (e *EventService) Confirm(c *gin.Context) {
	var cfm *ConfirmPost

	if err := c.ShouldBindJSON(&cfm); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf(constants.RequestFailed, err))
		log.Println(err)
	}

	log.Println(cfm)

	c.Data(http.StatusOK, "charset=utf8", []byte("ef5c48d6"))
}
