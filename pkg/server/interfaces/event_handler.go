package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
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
	case "APOD":
		apod, err := e.NasaApp.APOD()
		if err != nil {
			messageResponse.Message = constants.ServerErrorOccurred
			log.Printf(constants.RequestFailed, err)
			break
		}
		messageResponse.Message = apod.Title
		messageResponse.Message += "\n\n" + apod.Explanation
		messageResponse.Message += "\n\n" + "IMG: " + apod.HDUrl
		messageResponse.Message += "\n\n" + "Date: " + apod.Date
	case "FHAZ":
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
			messageResponse.Message = "Rover: " + photos.Photos[index].Rover.Name
			messageResponse.Message += "\n\nCamera: " + photos.Photos[index].Camera.FullName
			messageResponse.Message += "\n\n" + "IMG: " + photos.Photos[index].ImgSrc
			messageResponse.Message += "\n\n" + "Sol: " + fmt.Sprint(photos.Photos[index].Sol)
			messageResponse.Message += "\n\n" + "Date: " + photos.Photos[index].EarthDate
		}
	case "RHAZ":
		photos, err := e.NasaApp.GetMarsPhoto(&entities.MarsRoverPhotosRequest{
			Camera: constants.RoverCameraRHAZ,
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
			messageResponse.Message = "Rover: " + photos.Photos[index].Rover.Name
			messageResponse.Message += "\n\nCamera: " + photos.Photos[index].Camera.FullName
			messageResponse.Message += "\n\n" + "IMG: " + photos.Photos[index].ImgSrc
			messageResponse.Message += "\n\n" + "Sol: " + fmt.Sprint(photos.Photos[index].Sol)
			messageResponse.Message += "\n\n" + "Date: " + photos.Photos[index].EarthDate
		}
	case "MAST":
		photos, err := e.NasaApp.GetMarsPhoto(&entities.MarsRoverPhotosRequest{
			Camera: constants.RoverCameraMAST,
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
			messageResponse.Message = "Rover: " + photos.Photos[index].Rover.Name
			messageResponse.Message += "\n\nCamera: " + photos.Photos[index].Camera.FullName
			messageResponse.Message += "\n\n" + "IMG: " + photos.Photos[index].ImgSrc
			messageResponse.Message += "\n\n" + "Sol: " + fmt.Sprint(photos.Photos[index].Sol)
			messageResponse.Message += "\n\n" + "Date: " + photos.Photos[index].EarthDate
		}
	case "CHEMCAM":
		photos, err := e.NasaApp.GetMarsPhoto(&entities.MarsRoverPhotosRequest{
			Camera: constants.RoverCameraCHEMCAM,
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
			messageResponse.Message = "Rover: " + photos.Photos[index].Rover.Name
			messageResponse.Message += "\n\nCamera: " + photos.Photos[index].Camera.FullName
			messageResponse.Message += "\n\n" + "IMG: " + photos.Photos[index].ImgSrc
			messageResponse.Message += "\n\n" + "Sol: " + fmt.Sprint(photos.Photos[index].Sol)
			messageResponse.Message += "\n\n" + "Date: " + photos.Photos[index].EarthDate
		}
	case "Test #2.2":
		messageResponse.Message = "do - 6"
	default:
		messageResponse.Message = "I'm sorry, I don't understand you"
	}

	err := e.VkApp.SendMessage(messageResponse)

	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf(constants.RequestFailed, err))
		log.Printf(constants.RequestFailed, err)
		return
	}

	c.Data(http.StatusOK, "charset=utf8", []byte("ok"))
}

func (e *EventService) Confirm(c *gin.Context) {
	var cfm *ConfirmPost

	if err := c.ShouldBindJSON(&cfm); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf(constants.RequestFailed, err))
		log.Printf(constants.RequestFailed, err)
		return
	}

	log.Println(cfm)

	c.Data(http.StatusOK, "charset=utf8", []byte("ef5c48d6"))
}
