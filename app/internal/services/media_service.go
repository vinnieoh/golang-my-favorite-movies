package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/config"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/models"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/repositories"
	"github.com/vinnieoh/golang-my-favorite-movies/app/pkg/cache"
)

type MediaService struct {
	MediaRepo *repositories.MediaRepository
	RedisRepo *cache.RedisRepository
}

func NewMediaService(mediaRepo *repositories.MediaRepository, redisRepo *cache.RedisRepository) *MediaService {
	return &MediaService{
		MediaRepo: mediaRepo,
		RedisRepo: redisRepo,
	}
}

const baseURL = "https://api.themoviedb.org/3/"

func getHeaders() map[string]string {
	return map[string]string{
		"accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", config.Settings.APIMovie),
	}
}

func (s *MediaService) fetchFromAPI(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	headers := getHeaders()
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data, status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func (s *MediaService) getFromCacheOrAPI(cacheKey, url string) (interface{}, error) {
	cachedData, err := s.RedisRepo.Get(cacheKey)
	if err != nil {
		return nil, err
	}

	if cachedData != "" {
		var data interface{}
		if err := json.Unmarshal([]byte(cachedData), &data); err != nil {
			return nil, err
		}
		return data, nil
	}

	apiData, err := s.fetchFromAPI(url)
	if err != nil {
		return nil, err
	}

	if err := s.RedisRepo.Set(cacheKey, string(apiData), 24*time.Hour); err != nil {
		return nil, err
	}

	var data interface{}
	if err := json.Unmarshal(apiData, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (s *MediaService) GetTrendingAllWeekBR() (interface{}, error) {
	cacheKey := "trending_all_week_br"
	url := fmt.Sprintf("%strending/all/week?language=pt-BR", baseURL)
	return s.getFromCacheOrAPI(cacheKey, url)
}

func (s *MediaService) GetTrendingAllDayBR() (interface{}, error) {
	cacheKey := "trending_all_day_br"
	url := fmt.Sprintf("%strending/all/day?language=pt-BR", baseURL)
	return s.getFromCacheOrAPI(cacheKey, url)
}

func (s *MediaService) SearchContent(content string) (interface{}, error) {
	cacheKey := fmt.Sprintf("search_%s", content)
	url := fmt.Sprintf("%ssearch/multi?query=%s&include_adult=false&language=pt-BR&page=1", baseURL, content)
	return s.getFromCacheOrAPI(cacheKey, url)
}

func (s *MediaService) GetMovieByID(id string) (interface{}, error) {
	cacheKey := fmt.Sprintf("movie_%s", id)
	url := fmt.Sprintf("%smovie/%s?language=pt-BR", baseURL, id)
	return s.getFromCacheOrAPI(cacheKey, url)
}

func (s *MediaService) GetTVShowByID(id string) (interface{}, error) {
	cacheKey := fmt.Sprintf("tv_show_%s", id)
	url := fmt.Sprintf("%stv/%s?language=pt-BR", baseURL, id)
	return s.getFromCacheOrAPI(cacheKey, url)
}

// Movie methods
func (s *MediaService) GetMovies() ([]models.Movie, error) {
	return s.MediaRepo.GetMovies()
}

func (s *MediaService) CreateMovie(movie *models.Movie) error {
	return s.MediaRepo.CreateMovie(movie)
}

func (s *MediaService) UpdateMovie(movie *models.Movie) error {
	return s.MediaRepo.UpdateMovie(movie)
}

func (s *MediaService) DeleteMovie(id string) error {
	return s.MediaRepo.DeleteMovie(id)
}

// TVShow methods
func (s *MediaService) GetTVShows() ([]models.TVShow, error) {
	return s.MediaRepo.GetTVShows()
}

func (s *MediaService) CreateTVShow(tvShow *models.TVShow) error {
	return s.MediaRepo.CreateTVShow(tvShow)
}

func (s *MediaService) UpdateTVShow(tvShow *models.TVShow) error {
	return s.MediaRepo.UpdateTVShow(tvShow)
}

func (s *MediaService) DeleteTVShow(id string) error {
	return s.MediaRepo.DeleteTVShow(id)
}
