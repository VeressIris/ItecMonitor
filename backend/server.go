package main

import (
	"server/firestore"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	// init server:
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// get firestore client
	client, _ := firestore.GetFirestoreClient("itec2024revamp")
	defer client.Client.Close()

	app.Listen(":3000")
}
