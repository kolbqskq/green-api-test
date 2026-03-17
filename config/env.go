package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	InstanceID       string
	APITokenInstance string
	APIURL           string
	ChatID           string
}

func Init() *Config {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Не удалось прочитать .env")
	}

	return &Config{
		InstanceID:       getRequiredString("INSTANCE_ID"),
		APITokenInstance: getRequiredString("API_TOKEN"),
		APIURL:           getRequiredString("API_URL"),
		ChatID:           getRequiredString("CHAT_ID"),
	}
}

func getRequiredString(key string) string {
	str := os.Getenv(key)
	if str == "" {
		log.Fatalf("обязательная переменная окружения не задана: %s", key)
	}
	return str
}
