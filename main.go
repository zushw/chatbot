package main

import (
	"chatbot/controller"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	controller.RegisterChatRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
