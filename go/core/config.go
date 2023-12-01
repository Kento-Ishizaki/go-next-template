package core

import (
	"os"
)

const useMock = false

type Config struct {
	Port    string
	DB      DBConfig
	Env     string
	UseMock bool
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func NewConfig() *Config {
	return &Config{
		Port: os.Getenv("PORT"),
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
		},
		Env:     os.Getenv("ENV"),
		UseMock: useMock,
	}
}
