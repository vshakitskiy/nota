package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return errors.New("unable to load .env file")
	}

	return nil
}

func GetenvDefault(key, val string) string {
	res := os.Getenv(key)
	if res == "" {
		return val
	}

	return res
}
