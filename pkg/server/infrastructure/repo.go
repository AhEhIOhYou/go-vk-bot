package infrastructure

type Repos struct {
	Vk   VkRepo
	Nasa NasaRepo
}

func NewRepo(url string, accessToken string, version string) (*Repos, error) {

	return &Repos{
		Vk:   VkRepo{},
		Nasa: NasaRepo{},
	}, nil

}
