package config

import "time"

type Jwt struct {
	Exp time.Duration `yaml:"expiration"`
}

type Session struct {
	Exp time.Duration `yaml:"expiration"`
}

type Auth struct {
	Name         string   `yaml:"name"`
	Env          []string `yaml:"env"`
	ProtectedRPC []string `yaml:"protected_rpc"`
}

type Gateway struct {
	Name            string   `yaml:"name"`
	Env             []string `yaml:"env"`
	ProtectedRoutes []string `yaml:"protected_routes"`
}

type Snippet struct {
	Name string   `yaml:"name"`
	Env  []string `yaml:"env"`
}
