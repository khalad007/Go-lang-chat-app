package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	app.Use(cors.New())
	pusherClient := pusher.Client{
		AppID:   "1832921",
		Key:     "1817dc92d62fbdbbe132",
		Secret:  "f722d129140a7e64a942",
		Cluster: "ap2",
		Secure:  true,
	}
	// Define a route for the GET method on the root path '/'
	app.Post("/api/messages", func(c fiber.Ctx) error {

		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		pusherClient.Trigger("chat", "message", data)
		// Send a string response to the client
		return c.JSON([]string{})
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
