package main

import (
	"log"
	"net/http"

	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/config"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/handlers"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/migrations"
)

func main() {
    cfg := config.LoadConfig()
    router := handlers.SetupRouter()

    migrations.Migrate()  // Executa as migrações do banco de dados

    port := ":" + cfg.Server.Port
    log.Printf("Server is listening on port %s...", port)
    if err := http.ListenAndServe(port, router); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}