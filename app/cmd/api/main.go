package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/config"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/handlers"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/migrations"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/models"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/repositories"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/services"
)

func main() {
	cfg := config.LoadConfig()

	dsn := cfg.Database.URL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Defina a variável de ambiente JWT_SECRET
	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}

	// Auto-migrate the models
	db.AutoMigrate(&models.User{}, &models.Movie{}, &models.TVShow{}, &models.Comment{})

	// Initialize repositories
	userRepository := repositories.NewUserRepository(db)
	// movieRepository := repositories.NewMovieRepository(db)
	// tvShowRepository := repositories.NewTVShowRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepository)
	// movieService := services.NewMovieService(movieRepository)
	// tvShowService := services.NewTVShowService(tvShowRepository)

	// Initialize the router
	router := gin.Default()

	// Register routes
	handlers.RegisterRoutes(router, userService)

	migrations.Migrate() // Executa as migrações do banco de dados

	port := ":" + cfg.Server.Port
	log.Printf("Server is listening on port %s...", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
