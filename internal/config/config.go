package config

import (
	"log"
	"merch/pkg/database/postgres"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database postgres.Config
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Не удалось загрузить env")
	}

	dbConfig := postgres.Config{
		Addr:   os.Getenv("POSTGRES_HOST"),
		Port:   os.Getenv("POSTGRES_PORT"),
		User:   os.Getenv("POSTGRES_USER"),
		Passwd: os.Getenv("POSTGRES_PASSWORD"),
		DB:     os.Getenv("POSTGRES_DB"),
	}

	if dbConfig.User == "" || dbConfig.Passwd == "" {
		log.Fatal("POSTGRES_USER  POSTGRES_PASSWORD")
	}

	return &Config{Database: dbConfig}
}
