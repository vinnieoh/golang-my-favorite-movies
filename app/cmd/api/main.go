package main

import (
	"log"
	"net/http"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/config"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/handlers"
)

func main() {
    cfg := config.LoadConfig()
    router := handlers.SetupRouter()

    log.Printf("Server is listening on port %s...", cfg.Server.Port)
    if err := http.ListenAndServe(cfg.Server.Port, router); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}