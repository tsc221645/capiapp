package config

import (
	"log"
	"os"
)

var SecretKey []byte

func LoadEnv() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET no está definido")
	}
	SecretKey = []byte(secret)
	log.Println("JWT_SECRET cargado correctamente")
}
