package infrastructure

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/constants"
	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/entities"
	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/infrastructure/utils"

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

	_, err := http.NewRequest("GET", r.url+method, nil)
	if err != nil {
		return fmt.Errorf(constants.RequestCreationError, err)
	}

	keyboard := newKayboard()
	keyboardQuery, err := utils.KeyboardToQuery(&keyboard)
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

	values.Add("keyboard",keyboardQuery)

	log.Println("Query:")
	log.Println(values.Encode())

	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	return fmt.Errorf(constants.RequestFailed, err)
	// }

	// log.Println("Resp:")
	// log.Println(resp.Body)

	return nil
}
