package services

import (
    "github.com/vinnieoh/golang-my-favorite-movies/app/internal/models"
    "github.com/vinnieoh/golang-my-favorite-movies/app/internal/repositories"
)

func GetAllMovies() ([]models.Movie, error) {
    return repositories.GetAllMovies()
}
