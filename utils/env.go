package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	godotenv.Load()
}

func GetEnv(key string, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}

	return v
}
