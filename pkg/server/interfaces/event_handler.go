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
		return
	}

	// Object can be different types, so i use interface

	// There should be a switch case here, but so far we are working only with message_new, so there is no need

	var messageNew entities.MessageNew
	jsonData, _ := json.Marshal(data.Object)
	json.Unmarshal(jsonData, &messageNew)

	fmt.Println("Message text:")
	fmt.Println(messageNew.Message.Text)

	var messageResponse = &entities.MessageResponse{
		UserID: messageNew.Message.PeerID,
	}

	switch messageNew.Message.Text {
	case "Test #1":
		messageResponse.Message = "do - 1"
		apod, err := e.NasaApp.APOD()
		if err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Sprintf(constants.RequestFailed, err))
			log.Println(err)
			return
		}
		messageResponse.Message = apod.Title
		messageResponse.Message += "\n" + apod.Explanation
		messageResponse.Message += apod.HDUrl
	case "Test #2":
		messageResponse.Message = "do - 2"
	case "Test #3":
		messageResponse.Message = "do - 3"
	case "Test #4":
		messageResponse.Message = "do - 4"
	case "Test #2.1":
		messageResponse.Message = "do - 5"
	case "Test #2.2":
		messageResponse.Message = "do - 6"
	default:
		messageResponse.Message = "I'm sorry, I don't understand you"
	}

	err := e.VkApp.SendMessage(messageResponse)

	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf(constants.RequestFailed, err))
		log.Println(err)
		return
	}

	c.Data(http.StatusOK, "charset=utf8", []byte("ok"))
}

func (e *EventService) Confirm(c *gin.Context) {
	var cfm *ConfirmPost

	if err := c.ShouldBindJSON(&cfm); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf(constants.RequestFailed, err))
		log.Println(err)
		return
	}

	log.Println(cfm)

	c.Data(http.StatusOK, "charset=utf8", []byte("ef5c48d6"))
}
