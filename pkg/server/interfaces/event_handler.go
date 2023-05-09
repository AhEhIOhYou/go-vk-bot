package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/constants"
	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/entities"
	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/infrastructure"
	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/infrastructure/utils"
	"github.com/gin-gonic/gin"
)

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

func (e *EventService) NewVkEvent(c *gin.Context) {

	var data *entities.Event

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf(constants.RequestFailed, err))
		log.Println(err)
		return
	}

	if data.Secret != os.Getenv("VK_SECRET_KEY") {
		c.JSON(http.StatusInternalServerError, constants.SecretWordMissmatch)
		log.Print(constants.SecretWordMissmatch)
		return
	}

	if data.Type == "confirmation" {
		c.Data(http.StatusOK, "charset=utf8", []byte(os.Getenv("VK_API_CONFIRM")))
		return
	}

	// Object can be different types, so i use interface
	// There should be a switch case here, but so far we are working only with message_new, so there is no need

	var messageNew entities.MessageNew
	jsonData, _ := json.Marshal(data.Object)
	json.Unmarshal(jsonData, &messageNew)

	messageResponse := &entities.MessageResponse{
		UserID: messageNew.Message.PeerID,
	}

	// Message processing

	text := strings.ToLower(messageNew.Message.Text)

	// Switch-case to fine-tune actions
	switch text {
	case "apod":
		apod, err := e.NasaApp.APOD()
		if err != nil {
			messageResponse.Message = constants.ServerErrorOccurred
			log.Printf(constants.RequestFailed, err)
			break
		}
		messageResponse.Message = utils.PrepareAPODMessage(apod)
	case "fhaz":
		photos, err := e.NasaApp.GetMarsPhoto(&entities.MarsRoverPhotosRequest{
			Camera: constants.RoverCameraFHAZ,
			Sol:    rand.Intn(1000) + 600,
		})
		if err != nil {
			messageResponse.Message = constants.ServerErrorOccurred
			log.Printf(constants.RequestFailed, err)
			break
		}
		if len(photos.Photos) == 0 {
			messageResponse.Message = "Photo not found :("
		} else {
			index := rand.Intn(len(photos.Photos)) + 0
			messageResponse.Message = utils.PrepareMarsRoverPhotoMessage(&photos.Photos[index])
		}
	case "rhaz":
		photos, err := e.NasaApp.GetMarsPhoto(&entities.MarsRoverPhotosRequest{
			Camera: constants.RoverCameraRHAZ,
			Sol:    rand.Intn(1000) + 800,
		})
		if err != nil {
			messageResponse.Message = constants.ServerErrorOccurred
			log.Printf(constants.RequestFailed, err)
			break
		}
		if len(photos.Photos) == 0 {
			messageResponse.Message = "Photo not found :("
		} else {
			index := rand.Intn(len(photos.Photos)) + 0
			messageResponse.Message = utils.PrepareMarsRoverPhotoMessage(&photos.Photos[index])
		}
	case "mast":
		photos, err := e.NasaApp.GetMarsPhoto(&entities.MarsRoverPhotosRequest{
			Camera: constants.RoverCameraMAST,
			Sol:    rand.Intn(1000) + 800,
		})
		if err != nil {
			messageResponse.Message = constants.ServerErrorOccurred
			log.Printf(constants.RequestFailed, err)
			break
		}
		if len(photos.Photos) == 0 {
			messageResponse.Message = "Photo not found :("
		} else {
			index := rand.Intn(len(photos.Photos)) + 0
			messageResponse.Message = utils.PrepareMarsRoverPhotoMessage(&photos.Photos[index])
		}
	case "chemcam":
		photos, err := e.NasaApp.GetMarsPhoto(&entities.MarsRoverPhotosRequest{
			Camera: constants.RoverCameraCHEMCAM,
			Sol:    rand.Intn(1000) + 700,
		})
		if err != nil {
			messageResponse.Message = constants.ServerErrorOccurred
			log.Printf(constants.RequestFailed, err)
			break
		}
		if len(photos.Photos) == 0 {
			messageResponse.Message = "Photo not found :("
		} else {
			index := rand.Intn(len(photos.Photos)) + 0
			messageResponse.Message = utils.PrepareMarsRoverPhotoMessage(&photos.Photos[index])
		}
	case "test":
		messageResponse.Message = constants.TestSuccess
	case "commands":
		messageResponse.Message = constants.BotCommands
	case "description":
		messageResponse.Message = constants.BotDescription
	default:
		if utils.Contains(constants.Greetings, text) {
			messageResponse.Message = constants.BotDescription
			messageResponse.Message += "\n\n" + constants.BotCommands
		} else {
			messageResponse.Message = constants.BotUnknownCommandsMsg[rand.Intn(len(constants.BotUnknownCommandsMsg))]
		}
	}

	err := e.VkApp.SendMessage(messageResponse)

	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf(constants.RequestFailed, err))
		log.Printf(constants.RequestFailed, err)
		return
	}

	c.Data(http.StatusOK, "charset=utf8", []byte("ok"))
}
