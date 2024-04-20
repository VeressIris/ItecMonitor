package main

import (
	"fmt"
	"server/firestore"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// init server:
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// get firestore client
	client, _ := firestore.GetFirestoreClient("itec2024revamp")
	defer client.Client.Close()

	// add user to database endpoint
	app.Post("/addUser", func(c *fiber.Ctx) error {
		user := new(firestore.User)
		if err := c.BodyParser(user); err != nil {
			return err
		}
		client.WriteUserToDatabase(*user, user.UID)
		fmt.Printf("%s:%s added to database\n", user.Name, user.Email)
		return c.JSON("added user")
	})

	// add app to database endpoint
	app.Post("/addApp", func(c *fiber.Ctx) error {
		newApp := new(firestore.App)
		if err := c.BodyParser(newApp); err != nil {
			return err
		}
		client.WriteAppToDatabase(*newApp, newApp.Name)
		fmt.Printf("%s:%s added to database\n", newApp.Developer, newApp.Name)
		return c.JSON("added new app")
	})

	// get all apps from current user endpoint
	app.Post("/getDevApps", func(c *fiber.Ctx) error {
		dev := string(c.Body())
		devSplit := strings.Split(dev, "=")
		if len(devSplit) > 1 {
			dev = devSplit[1]
		}

		apps, err := client.GetDevApps(dev)
		if err != nil {
			return err
		}

		fmt.Print(dev + ": ")
		fmt.Println(apps)

		c.Set("Content-Type", "application/json")

		return c.JSON(apps)
	})

	app.Listen(":3000")
}
