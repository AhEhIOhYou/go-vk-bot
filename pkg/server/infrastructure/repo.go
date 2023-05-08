package infrastructure

import "os"

type Repos struct {
	Vk   VkRepo
	Nasa NasaRepo
}

func NewRepo() (*Repos, error) {

	return &Repos{
		Vk:   VkRepo{
			url: os.Getenv("VK_API_VERSION"),
			accessToken: os.Getenv("VK_API_VERSION"),
			version: os.Getenv("VK_API_VERSION"),
			methodNames: NewVkMethodNames(os.Getenv("VK_API_VERSION")),
		},
		Nasa: NasaRepo{
			
		},
	}, nil

}
