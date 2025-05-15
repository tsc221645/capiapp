package routes

import (
	"capiappproject/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/capyfacts", handlers.GetCapyFacts)
}
