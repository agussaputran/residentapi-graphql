package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Token global
var Token string

// GetEnvVar func
func GetEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Cannot load .env file")
	}
	return os.Getenv(key)
}
