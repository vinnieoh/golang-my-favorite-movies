package repositories

import (
    "github.com/vinnieoh/golang-my-favorite-movies/app/internal/models"
)

func GetAllMovies() ([]models.Movie, error) {
    // Aqui você implementaria a lógica para buscar os filmes do banco de dados
    return []models.Movie{
        {ID: 1, Title: "Inception"},
        {ID: 2, Title: "Interstellar"},
    }, nil
}
