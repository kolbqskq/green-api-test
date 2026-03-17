package config

import (
	"log"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	InstanceID       string `env:"INSTANCE_ID"`
	APITokenInstance string `env:"API_TOKEN"`
	APIURL           string `env:"API_URL"`
	ChatID           string `env:"CHAT_ID"`
}

func Init() *Config {
	var cfg Config

	godotenv.Load("../.env")

	if err := env.Parse(&cfg); err != nil {
		log.Fatal("Не удалось прочитать .env")
	}

	return &cfg
}

func getRequiredString(key string) string {
	str := os.Getenv(key)
	if str == "" {
		log.Fatalf("обязательная переменная окружения не задана: %s", key)
	}
	return str
}
