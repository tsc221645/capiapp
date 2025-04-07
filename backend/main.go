package main

import (
	"capiappproject/config"
	"capiappproject/database"
	"capiappproject/handlers"
	"capiappproject/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.LoadEnv()
	app := fiber.New()

	database.Connect()

	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Go backend using Fiber!")
	})

	app.Get("/profile", middleware.RequireAuth, func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")
		username := c.Locals("username")

		return c.JSON(fiber.Map{
			"user_id":  userID,
			"username": username,
		})
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // frontend
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Post("/signup", handlers.Register)
	app.Post("/login", handlers.Login)

	println("🚀 Server listening on http://localhost:3000")
	app.Listen(":3000")
}
