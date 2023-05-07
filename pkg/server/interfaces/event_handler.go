package interfaces

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/constants"
	"github.com/gin-gonic/gin"
)

type ConfirmPost struct {
	Type    string `json:"type"`
	GroupID int    `json:"group_id"`
}

type Event struct{}

func NewEvent() *Event {
	return &Event{}
}

func (e *Event) Index(c *gin.Context) {
	var cfm *ConfirmPost

	if err := c.ShouldBindJSON(&cfm); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf(constants.RequestFailed, err))
		log.Println()
	}

	log.Println(cfm)

	c.Data(http.StatusOK, "charset=utf8", []byte("8bb97950"))
}
