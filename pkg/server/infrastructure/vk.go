package infrastructure

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/constants"
	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/entities"

	qs "github.com/sonh/qs"
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
	arr := [][]entities.Button{
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
	}
	return entities.Keyboard{
		OneTime: false,
		Inline:  false,
		Buttons: arr,
	}
}

func (r *VkRepo) SendMessage(message *entities.MessageResponse) error {

	var method string = r.methodNames.sendMessage

	_, err := http.NewRequest("GET", r.url+method, nil)
	if err != nil {
		return fmt.Errorf(constants.RequestCreationError, err)
	}

	message.AccessToken = r.accessToken
	message.Version = r.version
	message.RandomID = rand.Intn(92233720368)
	message.Keyboard = newKayboard()

	encoder := qs.NewEncoder()
	test := message.Keyboard
	values, err := encoder.Values(test)
	if err != nil {
		return fmt.Errorf(constants.QueryCreationError, err)
	}

	// req.URL.RawQuery = values.Encode()

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
