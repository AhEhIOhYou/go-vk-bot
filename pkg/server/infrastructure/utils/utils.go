package utils

import (
	"fmt"

	"github.com/AhEhIOhYou/go-vk-bot/pkg/server/entities"
)

func PrepareAPODMessage(apod *entities.APOD) string {
	var resultStr string
	resultStr = apod.Title
	resultStr += "\n\n" + apod.Explanation
	resultStr += "\n\n" + "IMG: " + apod.HDUrl
	resultStr += "\n\n" + "Date: " + apod.Date
	return resultStr
}

func PrepareMarsRoverPhotoMessage(photo *entities.MarsRoverPhoto) string {
	var resultStr string
	resultStr = "Rover: " + photo.Rover.Name
	resultStr += "\n\nCamera: " + photo.Camera.FullName
	resultStr += "\n\n" + "IMG: " + photo.ImgSrc
	resultStr += "\n\n" + "Sol: " + fmt.Sprint(photo.Sol)
	resultStr += "\n\n" + "Date: " + photo.EarthDate
	return resultStr
}

// Amazing generics!
func Contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
