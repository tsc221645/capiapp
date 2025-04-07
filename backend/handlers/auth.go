package handlers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"capiappproject/database"
	"capiappproject/models"
)

// Handles user registration
func Register(c *fiber.Ctx) error {
	type Request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error hashing password"})
	}

	user := models.User{
		Username:     body.Username,
		Email:        body.Email,
		PasswordHash: string(hashedPassword),
	}

	// Check if username or email already exists
	// If user exists, return an error
	var existingUser models.User
	if err := database.DB.
		Where("username = ? OR email = ?", body.Username, body.Email).
		First(&existingUser).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid username or email!"})
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Could not create user"})
	}

	return c.JSON(fiber.Map{"message": "User registered successfully"}) // Return a succes message when correctly registered
}

// Handles user login
func Login(c *fiber.Ctx) error {
	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	if err := database.DB.Where("username = ?", body.Username).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Error finding user"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Incorrect password"})
	}

	return c.JSON(fiber.Map{
		"message":  "Successfullyt logged in",
		"username": user.Username,
		"email":    user.Email,
	})
}
