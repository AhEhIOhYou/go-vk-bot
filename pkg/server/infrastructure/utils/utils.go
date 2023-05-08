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

	return url.QueryEscape(string(jsonBytes)), nil
}
