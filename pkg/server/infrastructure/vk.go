package infrastructure

import (
	"encoding/json"
	"fmt"
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
			},
		},
	}
}

func (r *VkRepo) SendMessage(message *entities.MessageResponse) error {

	var method string = r.methodNames.sendMessage

	req, err := http.NewRequest("GET", r.url+method, nil)
	if err != nil {
		return fmt.Errorf(constants.RequestCreationError, err)
	}

	keyboard := newKayboard()
	keyboardJson, err := json.Marshal(keyboard)
	if err != nil {
		return fmt.Errorf(constants.QueryCreationError, err)
	}

	message.AccessToken = r.accessToken
	message.Version = r.version
	message.RandomID = rand.Intn(92233720368)
	message.Message = "ahahahhahahah"

	values, err := qs.Values(message)
	if err != nil {
		return fmt.Errorf(constants.QueryCreationError, err)
	}

	values.Set("keyboard", string(keyboardJson))

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf(constants.RequestFailed, err)
	}

	return nil
}
