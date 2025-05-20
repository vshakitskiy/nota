package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(path string) error {
	err := godotenv.Load(path)
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
