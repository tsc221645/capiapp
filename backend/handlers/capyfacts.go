// handlers/capyfacts.go
package handlers

import (
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func GetCapyFacts(c *fiber.Ctx) error {
	req, err := http.NewRequest("GET", "https://api.api-ninjas.com/v1/animals?name=capybara", nil)
	if err != nil {
		println("❌ Error creando request:", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Error creando solicitud")
	}

	req.Header.Add("X-Api-Key", os.Getenv("NINJA_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		println("❌ Error haciendo request:", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Error obteniendo datos de la API")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		println("❌ Error leyendo respuesta:", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString("Error leyendo respuesta de la API")
	}

	println("✅ Datos recibidos desde API externa")
	c.Set("Content-Type", "application/json")
	return c.Send(body)
}
