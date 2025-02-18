package config

import (
	"os"

	"github.com/backent/ai-golang/helpers"
	"github.com/joho/godotenv"
)

var geminiAPI string

func init() {
	err := godotenv.Load()
	helpers.PanicIfError(err)
	if geminiAPI == "" {
		geminiAPI = os.Getenv("GEMINI_API_KEY")
	}
}

func GetGeminiAPIKey() string {
	return geminiAPI
}
