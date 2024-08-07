package main

import (
	"log"

	"apu/internal/api"
	"apu/internal/http/gin"
)

func main() {
	//s := book.NewService(postgres.New())
	s := initHandler()

	h := gin.Handlers(s)
	//h := fiber.Handlers(s)
	//hertz.Handlers(s)

	err := api.Start("8080", h)
	if err != nil {
		log.Fatalf("failed to start api server: %v", err)
	}
}
