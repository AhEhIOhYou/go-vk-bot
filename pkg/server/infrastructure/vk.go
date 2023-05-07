package infrastructure

import (
	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/entities"
)

type VkRepo struct {
	url         string
	accessToken string
	version     string
}

func (r *VkRepo) SendMessage(message *entities.MessageResponse) error {

	return nil
}
