package infrastructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/constants"
	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/entities"

	qs "github.com/google/go-querystring/query"
)

type VkMethodNames struct {
	sendMessage string
}

type VkRepo struct {
	url         string
	accessToken string
	version     string
	methodNames *VkMethodNames
}

func NewVkMethodNames(sendMessage string) *VkMethodNames {
	return &VkMethodNames{
		sendMessage: sendMessage,
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
						Label: "FHAZ",
					},
				},
				{
					Color: "primary",
					Action: entities.ButtonAction{
						Type:  "text",
						Label: "RHAZ",
					},
				},
				{
					Color: "primary",
					Action: entities.ButtonAction{
						Type:  "text",
						Label: "MAST",
					},
				},
				{
					Color: "primary",
					Action: entities.ButtonAction{
						Type:  "text",
						Label: "CHEMCAM",
					},
				},
			},
			{
				{
					Color: "secondary",
					Action: entities.ButtonAction{
						Type:  "text",
						Label: "APOD",
					},
				},
				{
					Color: "positive",
					Action: entities.ButtonAction{
						Type:  "text",
						Label: "Test",
					},
				},
			},
			{
				{
					Color: "primary",
					Action: entities.ButtonAction{
						Type:  "text",
						Label: "Description",
					},
				},
			},
			{
				{
					Color: "secondary",
					Action: entities.ButtonAction{
						Type:  "text",
						Label: "Commands",
					},
				},
			},
		},
	}
}

var keyboard = newKayboard()

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

// Unused
// I tried uploading photos, but realized that the VK itself downloads photos by the link
func (r *VkRepo) GetMessageUploadServer(uploadReq *entities.MessageUploadServerRequest) (*entities.VkResponse, error) {

	var method string = "photos.getMessagesUploadServer"

	// Prepare values
	uploadReq.AccessToken = r.accessToken
	uploadReq.Version = r.version

	values, err := qs.Values(uploadReq)
	if err != nil {
		return nil, fmt.Errorf(constants.QueryCreationError, err)
	}

	// Prepare request
	req, err := http.NewRequest("GET", r.url+method, nil)
	if err != nil {
		return nil, fmt.Errorf(constants.RequestCreationError, err)
	}

	req.URL.RawQuery = values.Encode()

	// Execute request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(constants.RequestFailed, err)
	}

	defer resp.Body.Close()

	var uploadResp *entities.VkResponse

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf(constants.DecodingJSONError, err)
	}

	err = json.Unmarshal(body, &uploadResp)
	if err != nil {
		return nil, fmt.Errorf(constants.DecodingJSONError, err)
	}

	return uploadResp, nil
}
