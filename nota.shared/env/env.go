package env

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(envs []string) error {
	err := godotenv.Load("../.env")
	if err != nil {
		return errors.New("unable to load .env file from root path. See .env.example")
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		return errors.New("APP_ENV is not set. See .env.example")
	}

	for _, env := range envs {
		switch env {
		case "jwt":
			err = loadJwt()
		case "postgres":
			err = loadPostgres()
		case "otel":
			err = loadOtelCollector()
		case "auth":
			err = loadAuth()
		case "gateway":
			err = loadGateway()
		case "snippet":
			err = loadSnippet()
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func GetAppEnv() string {
	return os.Getenv("APP_ENV")
}

func GetJwtSecret() string {
	return os.Getenv("JWT_SECRET")
}

func GetPostgresUser() string {
	return os.Getenv("POSTGRES_USER")
}

func GetPostgresPassword() string {
	return os.Getenv("POSTGRES_PASSWORD")
}

func GetPostgresHost() string {
	return os.Getenv("POSTGRES_HOST")
}

func GetPostgresPort() string {
	return os.Getenv("POSTGRES_PORT")
}

func GetPostgresDB() string {
	return os.Getenv("POSTGRES_DB")
}

func GetOtelCollector() string {
	return os.Getenv("OTEL_COLLECTOR")
}

func GetAuthHost() string {
	return os.Getenv("AUTH_HOST")
}

func GetAuthPort() string {
	return os.Getenv("AUTH_PORT")
}

func GetGatewayHost() string {
	return os.Getenv("GATEWAY_HOST")
}

func GetGatewayPort() string {
	return os.Getenv("GATEWAY_PORT")
}

func GetSnippetHost() string {
	return os.Getenv("SNIPPET_HOST")
}

func GetSnippetPort() string {
	return os.Getenv("SNIPPET_PORT")
}

func loadJwt() error {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return errors.New("JWT_SECRET is not set. See .env.example")
	}

	return nil
}

func loadPostgres() error {
	postgresUser := os.Getenv("POSTGRES_USER")
	if postgresUser == "" {
		return errors.New("POSTGRES_USER is not set. See .env.example")
	}

	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	if postgresPassword == "" {
		return errors.New("POSTGRES_PASSWORD is not set. See .env.example")
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost == "" {
		return errors.New("POSTGRES_HOST is not set. See .env.example")
	}

	postgresPort := os.Getenv("POSTGRES_PORT")
	if postgresPort == "" {
		return errors.New("POSTGRES_PORT is not set. See .env.example")
	}

	postgresDB := os.Getenv("POSTGRES_DB")
	if postgresDB == "" {
		return errors.New("POSTGRES_DB is not set. See .env.example")
	}

	return nil
}

func loadOtelCollector() error {
	otelCollector := os.Getenv("OTEL_COLLECTOR")
	if otelCollector == "" {
		return errors.New("OTEL_COLLECTOR is not set. See .env.example")
	}

	return nil
}

func loadAuth() error {
	authHost := os.Getenv("AUTH_HOST")
	if authHost == "" {
		return errors.New("AUTH_HOST is not set. See .env.example")
	}

	authPort := os.Getenv("AUTH_PORT")
	if authPort == "" {
		return errors.New("AUTH_PORT is not set. See .env.example")
	}

	return nil
}

func loadSnippet() error {
	snippetHost := os.Getenv("SNIPPET_HOST")
	if snippetHost == "" {
		return errors.New("SNIPPET_HOST is not set. See .env.example")
	}

	snippetPort := os.Getenv("SNIPPET_PORT")
	if snippetPort == "" {
		return errors.New("SNIPPET_PORT is not set. See .env.example")
	}

	return nil
}

func loadGateway() error {
	gatewayHost := os.Getenv("GATEWAY_HOST")
	if gatewayHost == "" {
		return errors.New("GATEWAY_HOST is not set. See .env.example")
	}

	gatewayPort := os.Getenv("GATEWAY_PORT")
	if gatewayPort == "" {
		return errors.New("GATEWAY_PORT is not set. See .env.example")
	}

	return nil
}
