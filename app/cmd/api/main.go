package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/config"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/handlers"
	//"github.com/vinnieoh/golang-my-favorite-movies/app/internal/migrations"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/models"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/repositories"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/services"
	"github.com/vinnieoh/golang-my-favorite-movies/app/pkg/cache"
)

func main() {
	cfg := config.LoadConfig()

	dsn := cfg.Database.URL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Verifique se a variável de ambiente JWT_SECRET está configurada
	if cfg.JWT.Secret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}

	// Auto-migrate the models
	err = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if err != nil {
		log.Fatalf("Failed to create extension uuid-ossp: %v", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Movie{}, &models.TVShow{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize Redis repository
	redisRepo := cache.NewRedisRepository(
		cfg.Redis.Host+":"+cfg.Redis.Port,
		"", // Assuming no password for Redis
		0,  // Using default DB
	)

	// Initialize repositories
	userRepository := repositories.NewUserRepository(db)
	mediaRepository := repositories.NewMediaRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepository)
	mediaService := services.NewMediaService(mediaRepository, redisRepo)

	// Initialize the router
	router := gin.Default()

	// Register routes
	handlers.RegisterRoutes(router, userService, mediaService, cfg.JWT.Secret)

	// Executa as migrações do banco de dados
	//migrations.Migrate()

	port := ":" + cfg.Server.Port
	log.Printf("Server is listening on port %s...", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
