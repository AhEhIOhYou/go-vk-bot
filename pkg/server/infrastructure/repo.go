package infrastructure

import "os"

type Repos struct {
	Vk   VkRepo
	Nasa NasaRepo
}

func NewRepo() (*Repos, error) {
	return &Repos{
		Vk: VkRepo{
			url:         os.Getenv("VK_API_URL"),
			accessToken: os.Getenv("VK_API_KEY"),
			version:     os.Getenv("VK_API_VERSION"),
			methodNames: NewVkMethodNames(
				os.Getenv("VK_API_METHOD_MESSAGE_SEND"),
			),
		},
		Nasa: NasaRepo{
			url:         os.Getenv("NASA_API_URL"),
			accessToken: os.Getenv("NASA_API_KEY"),
			marsUrl:     os.Getenv("NASA_MARS_WEATHER_URL"),
			methodNames: NewNasaMethodNames(
				os.Getenv("NASA_API_METHOD_APOD"),
				os.Getenv("NASA_API_METHOD_MARS_PHOTO"),
			),
		},
	}, nil

}
