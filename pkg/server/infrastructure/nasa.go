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
	apod           string
	marsRoverPhoto string
}

type NasaRepo struct {
	url         string
	marsUrl     string
	accessToken string
	methodNames *NasaMethodNames
}

func NewNasaMethodNames(apod, marsRoverPhoto string) *NasaMethodNames {
	return &NasaMethodNames{
		apod:           apod,
		marsRoverPhoto: marsRoverPhoto,
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

	// Prepare request
	req, err := http.NewRequest("GET", r.url+method, nil)
	if err != nil {
		return nil, fmt.Errorf(constants.RequestCreationError, err)
	}

	req.URL.RawQuery = values.Encode()

	// Execute request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(constants.RequestFailed, err)
	}

	defer resp.Body.Close()

	var apod []entities.APOD

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf(constants.DecodingJSONError, err)
	}

	err = json.Unmarshal(body, &apod)
	if err != nil {
		return nil, fmt.Errorf(constants.DecodingJSONError, err)
	}

	return &apod[0], nil
}

func (r *NasaRepo) GetMarsPhoto(roverPhotoReq *entities.MarsRoverPhotosRequest) (*entities.MarsRoverPhotots, error) {

	var method string = r.methodNames.marsRoverPhoto

	roverPhotoReq.ApiKey = r.accessToken

	log.Println(roverPhotoReq)

	values, err := qs.Values(roverPhotoReq)
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

	log.Println("REQ::")
	log.Println(resp.Request.URL)

	var photos entities.MarsRoverPhotots

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf(constants.DecodingJSONError, err)
	}

	err = json.Unmarshal(body, &photos)
	if err != nil {
		return nil, fmt.Errorf(constants.DecodingJSONError, err)
	}

	return &photos, nil
}
