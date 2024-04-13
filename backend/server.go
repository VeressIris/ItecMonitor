package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	// server stuff:
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Init firebase
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "itec2024revamp"}
	firebaseApp, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalf("Failed to init firebase app: %v", err)
	}

	// init firestore
	client, err := firebaseApp.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to init firestore: %v", err)
	}
	defer client.Close()

	app.Listen(":3000")
}
