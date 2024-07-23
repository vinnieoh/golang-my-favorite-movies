package repositories

import (
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/models"
	"gorm.io/gorm"
)

type MediaRepository struct {
	DB *gorm.DB
}

func NewMediaRepository(db *gorm.DB) *MediaRepository {
	return &MediaRepository{DB: db}
}

// Movie methods
func (r *MediaRepository) GetMovies() ([]models.Movie, error) {
	var movies []models.Movie
	err := r.DB.Find(&movies).Error
	return movies, err
}

func (r *MediaRepository) GetMovieByID(id string) (*models.Movie, error) {
	var movie models.Movie
	err := r.DB.First(&movie, "id = ?", id).Error
	return &movie, err
}

func (r *MediaRepository) CreateMovie(movie *models.Movie) error {
	return r.DB.Create(movie).Error
}

func (r *MediaRepository) UpdateMovie(movie *models.Movie) error {
	return r.DB.Save(movie).Error
}

func (r *MediaRepository) DeleteMovie(id string) error {
	return r.DB.Delete(&models.Movie{}, "id = ?", id).Error
}

// TVShow methods
func (r *MediaRepository) GetTVShows() ([]models.TVShow, error) {
	var tvShows []models.TVShow
	err := r.DB.Find(&tvShows).Error
	return tvShows, err
}

func (r *MediaRepository) GetTVShowByID(id string) (*models.TVShow, error) {
	var tvShow models.TVShow
	err := r.DB.First(&tvShow, "id = ?", id).Error
	return &tvShow, err
}

func (r *MediaRepository) CreateTVShow(tvShow *models.TVShow) error {
	return r.DB.Create(tvShow).Error
}

func (r *MediaRepository) UpdateTVShow(tvShow *models.TVShow) error {
	return r.DB.Save(tvShow).Error
}

func (r *MediaRepository) DeleteTVShow(id string) error {
	return r.DB.Delete(&models.TVShow{}, "id = ?", id).Error
}
