package infrastructure

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/constants"
	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/entities"

	querystring "github.com/google/go-querystring/query"
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

func (r *VkRepo) SendMessage(message *entities.MessageResponse) error {

	var method string = r.methodNames.sendMessage

	req, err := http.NewRequest("GET", r.url+method, nil)
	if err != nil {
		return fmt.Errorf(constants.RequestCreationError, err)
	}

	message.AccessToken = r.accessToken
	message.Version = r.version
	message.RandomID = rand.Intn(92233720368)

	vals, err := querystring.Values(message)
	if err != nil {
		return fmt.Errorf(constants.QueryCreationError, err)
	}

	req.URL.RawQuery = vals.Encode()

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf(constants.RequestFailed, err)
	}

	return nil
}
