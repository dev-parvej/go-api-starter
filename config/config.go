package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Get(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}
