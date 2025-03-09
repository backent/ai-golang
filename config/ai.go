package config

import (
	"log"
	"os"
	"runtime/debug"

	"github.com/joho/godotenv"
)

var geminiAPI string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("ERROR: %v\n%s", err, debug.Stack())
	}

	if geminiAPI == "" {
		geminiAPI = os.Getenv("GEMINI_API_KEY")
	}
}

func GetGeminiAPIKey() string {
	return geminiAPI
}
