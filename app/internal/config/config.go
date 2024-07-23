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
    Redis struct {
        Host string
        Port string
        DB   string
    }
    JWT struct {
        Secret    string
        Algorithm string
    }
    APIMovie string
}

var Settings *Config

func LoadConfig() *Config {
    if err := godotenv.Load("./dotenv_files/.env"); err != nil {
        log.Printf("Error loading .env file")
    }

    cfg := &Config{}
    cfg.Server.Port = os.Getenv("SERVER_PORT")
    cfg.Database.URL = os.Getenv("DB_URL")
    cfg.Redis.Host = os.Getenv("HOST_REDIS")
    cfg.Redis.Port = os.Getenv("PORT_REDIS")
    cfg.Redis.DB = os.Getenv("DB_REDIS")
    cfg.JWT.Secret = os.Getenv("JWT_SECRET")
    cfg.JWT.Algorithm = os.Getenv("ALGORITHM")
    cfg.APIMovie = os.Getenv("API_MOVIE")

    if cfg.APIMovie == "" {
        log.Fatal("API_MOVIE environment variable is required")
    }

    Settings = cfg
    return cfg
}
