package utils

import (
	"encoding/json"
	"net/url"

	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/entities"
)

func KeyboardToQuery(keyboard *entities.Keyboard) (string, error) {
	jsonBytes, err := json.Marshal(keyboard)
	if err != nil {
		return "", err
	}
	jsonStr := string(jsonBytes)

	encodedStr := url.QueryEscape(jsonStr)
	return encodedStr, nil
}
