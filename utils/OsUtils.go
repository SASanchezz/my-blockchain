package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// init is invoked before main()
func InitEnvVariables() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
