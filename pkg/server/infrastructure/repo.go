package infrastructure

import (
	"os"

	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/constants"
)

type Repos struct {
	Vk   VkRepo
	Nasa NasaRepo
}

func NewRepo() (*Repos, error) {
	return &Repos{
		Vk: VkRepo{
			url:         constants.VkApiUrl,
			accessToken: os.Getenv("VK_API_KEY"),
			version:     constants.VkApiVersion,
			methodNames: NewVkMethodNames(
				constants.VkApiMethodMessageSend,
			),
		},
		Nasa: NasaRepo{
			url:         constants.NasaApiUrl,
			accessToken: os.Getenv("NASA_API_KEY"),
			methodNames: NewNasaMethodNames(
				constants.NasaApiMethodApod,
				constants.NasaApiMethodMarsPhoto,
			),
		},
	}, nil

}
