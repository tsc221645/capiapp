package main

import (
	"capiappproject/database"
	"capiappproject/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Connect()

	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Go backend using Fiber! 🚀")
	})

	app.Post("/register", handlers.Register)

	println("🚀 Server listening on http://localhost:3000")
	app.Listen(":3000")
}
