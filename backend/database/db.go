package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"capiappproject/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("[retry %d] No se pudo conectar a la base de datos: %v", i+1, err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatalf("Fallo final al conectar a la base de datos: %v", err)
	}
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("No se pudo migrar el modelo User: %v", err)
	}
}
