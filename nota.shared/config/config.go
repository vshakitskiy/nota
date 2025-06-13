package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type jwt struct {
	Jwt Jwt `yaml:"jwt"`
}

type session struct {
	Session Session `yaml:"session"`
}

type auth struct {
	Auth Auth `yaml:"auth"`
}

type gateway struct {
	Gateway Gateway `yaml:"gateway"`
}

type snippet struct {
	Snippet Snippet `yaml:"snippet"`
}

func LoadJwt() (*Jwt, error) {
	cfg, err := readConfig()
	if err != nil {
		return nil, err
	}

	jwt := &jwt{}
	err = yaml.Unmarshal(cfg, &jwt)
	if err != nil {
		return nil, errors.New("failed to unmarshal config file")
	}

	return &jwt.Jwt, nil
}

func LoadSession() (*Session, error) {
	cfg, err := readConfig()
	if err != nil {
		return nil, err
	}

	session := &session{}
	err = yaml.Unmarshal(cfg, &session)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("failed to unmarshal config file")
	}

	return &session.Session, nil
}

func LoadAuth() (*Auth, error) {
	cfg, err := readConfig()
	if err != nil {
		return nil, err
	}

	auth := &auth{}
	err = yaml.Unmarshal(cfg, &auth)
	if err != nil {
		return nil, errors.New("failed to unmarshal config file")
	}

	return &auth.Auth, nil
}

func LoadGateway() (*Gateway, error) {
	cfg, err := readConfig()
	if err != nil {
		return nil, err
	}

	gateway := &gateway{}
	err = yaml.Unmarshal(cfg, &gateway)
	if err != nil {
		return nil, errors.New("failed to unmarshal config file")
	}

	return &gateway.Gateway, nil
}

func LoadSnippet() (*Snippet, error) {
	cfg, err := readConfig()
	if err != nil {
		return nil, err
	}

	snippet := &snippet{}
	err = yaml.Unmarshal(cfg, &snippet)
	if err != nil {
		return nil, errors.New("failed to unmarshal config file")
	}

	return &snippet.Snippet, nil
}

func readConfig() ([]byte, error) {
	f, err := os.ReadFile("../config/dev.yaml")
	if err != nil {
		return nil, errors.New("failed to read config file from root config folder")
	}

	return f, nil
}
