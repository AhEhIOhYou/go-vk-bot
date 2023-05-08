package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/constants"
	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/entities"
	qs "github.com/google/go-querystring/query"
)

type NasaMethodNames struct {
	apod string
}

type NasaRepo struct {
	url         string
	accessToken string
	methodNames *NasaMethodNames
}

func NewNasaMethodNames(apod string) *NasaMethodNames {
	return &NasaMethodNames{
		apod: apod,
	}
}

func (r *NasaRepo) APOD() (*entities.APOD, error) {

	var method string = r.methodNames.apod

	// Prepare values
	apodRequest := &entities.APODRequset{
		ApiKey: r.accessToken,
		Date:   "today",
		Count:  1,
	}

	values, err := qs.Values(apodRequest)
	if err != nil {
		return nil, fmt.Errorf(constants.QueryCreationError, err)
	}

	// Prepare request
	req, err := http.NewRequest("GET", r.url+method, nil)
	if err != nil {
		return nil, fmt.Errorf(constants.RequestCreationError, err)
	}

	req.URL.RawQuery = values.Encode()

	log.Println(req.URL.RawQuery)

	// Execute request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(constants.RequestFailed, err)
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var apod *entities.APOD
	err = decoder.Decode(&apod)
	if err != nil {
		return nil, fmt.Errorf(constants.DecodingJSONError, err)
	}

	log.Println("NASA Result:")
	log.Println(apod)

	return apod, nil
}
