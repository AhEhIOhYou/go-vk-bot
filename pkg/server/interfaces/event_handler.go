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

	fmt.Println(data.Type)

	// Object can be different types, so i use interface

	var messageNew entities.MessageNew
	jsonData, _ := json.Marshal(data.Object)
	json.Unmarshal(jsonData, &messageNew)

	fmt.Println("Message text:")
	fmt.Println(messageNew.Message.Text)

	e.VkApp.SendMessage(&entities.MessageResponse{
		Message: "test-response-6",
		UserID:  messageNew.Message.PeerID,
	})

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
