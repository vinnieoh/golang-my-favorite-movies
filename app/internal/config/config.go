package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    Server struct {
        Port string
    }
    Database struct {
        URL string
    }
}

func LoadConfig() *Config {
    if err := godotenv.Load(); err != nil {
        log.Printf("Error loading .env file")
    }

    cfg := &Config{}
    cfg.Server.Port = os.Getenv("SERVER_PORT")
    cfg.Database.URL = os.Getenv("DB_URL")

    return cfg
}
