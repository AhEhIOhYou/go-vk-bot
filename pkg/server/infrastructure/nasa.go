package infrastructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		Count:  1,
	}

	values, err := qs.Values(apodRequest)
	if err != nil {
		return nil, fmt.Errorf(constants.QueryCreationError, err)
	}

	// Execute request
	path := r.url+method + "?" + values.Encode()
	log.Println(path)
	resp, err := http.Get(path)
	if err != nil {
		return nil, fmt.Errorf(constants.RequestCreationError, err)
	}

	log.Println(resp.Request.URL)

	defer resp.Body.Close()

	var apod entities.APOD

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf(constants.DecodingJSONError, err)
	}

	err = json.Unmarshal(body, &apod)
	if err != nil {
		return nil, fmt.Errorf(constants.DecodingJSONError, err)
	}

	log.Println("NASA Status:")
	log.Println(resp.Status)
	log.Println("NASA Result:")
	log.Println(apod)

	return &apod, nil
}
