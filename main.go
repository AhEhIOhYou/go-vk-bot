package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello vk api")

	router := gin.Default()

	log.Fatal(router.Run(":8094"))
}
