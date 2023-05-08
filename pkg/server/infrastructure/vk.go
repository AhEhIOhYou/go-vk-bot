package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/constants"
	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/entities"

	qs "github.com/google/go-querystring/query"
)

type VkMethodNames struct {
	sendMessage            string
	sendMessageEventAnswer string
}

type VkRepo struct {
	url         string
	accessToken string
	version     string
	methodNames *VkMethodNames
}

func NewVkMethodNames(sendMessage, sendMessageEventAnswer string) *VkMethodNames {
	return &VkMethodNames{
		sendMessage:            sendMessage,
		sendMessageEventAnswer: sendMessageEventAnswer,
	}
}

func NewVkRepo(url, access_token, version string, methodNames *VkMethodNames) *VkRepo {
	return &VkRepo{
		url:         url,
		accessToken: access_token,
		version:     version,
		methodNames: methodNames,
	}
}

func newKayboard() entities.Keyboard {
	return entities.Keyboard{
		OneTime: false,
		Inline:  false,
		Buttons: [][]entities.Button{
			{
				{
					Color: "primary",
					Action: entities.ButtonAction{
						Type:  "text",
						Label: "Test #1",
					},
				},
				{
					Color: "primary",
					Action: entities.ButtonAction{
						Type:  "text",
						Label: "Test #2",
					},
				},
				{
					Color: "negative",
					Action: entities.ButtonAction{
						Type:    "callback",
						Label:   "Test #3",
						Payload: "btn#3",
					},
				},
			},
		},
	}
}

func newPopUpEvent() entities.EventData {
	return entities.EventData{
		Type: "show_snackbar",
		Text: "Hello there, ahehiohyou!",
	}
}

var keyboard = newKayboard()
var popUpEvent = newPopUpEvent()

func (r *VkRepo) SendMessage(message *entities.MessageResponse) error {

	var method string = r.methodNames.sendMessage

	// Prepare values
	keyboardJson, err := json.Marshal(keyboard)
	if err != nil {
		return fmt.Errorf(constants.QueryCreationError, err)
	}

	message.AccessToken = r.accessToken
	message.Version = r.version
	message.RandomID = rand.Intn(92233720368)

	values, err := qs.Values(message)
	if err != nil {
		return fmt.Errorf(constants.QueryCreationError, err)
	}

	values.Set("keyboard", string(keyboardJson))

	// Prepare request
	req, err := http.NewRequest("GET", r.url+method, nil)
	if err != nil {
		return fmt.Errorf(constants.RequestCreationError, err)
	}

	req.URL.RawQuery = values.Encode()

	// Execute request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf(constants.RequestFailed, err)
	}

	log.Println("Resp:")
	log.Println(resp.Status)

	return nil
}

func (r *VkRepo) SendEvent(eventMessage *entities.EventResponse) error {

	var method string = r.methodNames.sendMessageEventAnswer

	// Prepare values
	eventDataJson, err := json.Marshal(popUpEvent)
	if err != nil {
		return fmt.Errorf(constants.QueryCreationError, err)
	}

	eventMessage.AccessToken = r.accessToken
	eventMessage.Version = r.version

	values, err := qs.Values(eventMessage)
	if err != nil {
		return fmt.Errorf(constants.QueryCreationError, err)
	}

	values.Set("event_data", string(eventDataJson))

	// Prepare request
	req, err := http.NewRequest("GET", r.url+method, nil)
	if err != nil {
		return fmt.Errorf(constants.RequestCreationError, err)
	}

	req.URL.RawQuery = values.Encode()

	// Execute request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf(constants.RequestFailed, err)
	}

	log.Println("Resp:")
	log.Println(resp.Status)

	return nil
}
