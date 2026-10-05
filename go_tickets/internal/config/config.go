package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string
	Dsn  string
	JWTSecret string
}

func LoadEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	cfg := &Config{
		Port: os.Getenv("PORT"),
		Dsn:  os.Getenv("DSN"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
	if cfg.Port == "" {
		log.Fatal("PORT is not set")
	}
	if cfg.Dsn == "" {
		log.Fatal("DSN is not set")
	}

	return cfg
}
