package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	server := fiber.New()
	log.Fatal(server.Listen(":8080"))
}
