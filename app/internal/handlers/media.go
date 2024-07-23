package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "github.com/vinnieoh/golang-my-favorite-movies/app/internal/models"
    "github.com/vinnieoh/golang-my-favorite-movies/app/internal/services"
)

type MediaHandler struct {
    MediaService *services.MediaService
}

func NewMediaHandler(mediaService *services.MediaService) *MediaHandler {
    return &MediaHandler{
        MediaService: mediaService,
    }
}

// Movie handlers
func (h *MediaHandler) GetMovies(c *gin.Context) {
    movies, err := h.MediaService.GetMovies()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, movies)
}

func (h *MediaHandler) GetMovie(c *gin.Context) {
    id := c.Param("id")
    movie, err := h.MediaService.GetMovieByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
        return
    }
    c.JSON(http.StatusOK, movie)
}

func (h *MediaHandler) CreateMovie(c *gin.Context) {
    var movie models.Movie
    if err := c.ShouldBindJSON(&movie); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Captura o usuário autenticado
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
        return
    }

    // Define o ID do usuário
    movie.UserID = userID.(uuid.UUID)

    if err := h.MediaService.CreateMovie(&movie); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, movie)
}

func (h *MediaHandler) UpdateMovie(c *gin.Context) {
    id := c.Param("id")
    var movie models.Movie
    data, err := h.MediaService.GetMovieByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
        return
    }

    // Faça a assertiva de tipo
    movie, ok := data.(models.Movie)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid movie data"})
        return
    }

    if err := c.ShouldBindJSON(&movie); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    movie.ID = uuid.MustParse(id) // Garantir que o ID do filme seja mantido
    if err := h.MediaService.UpdateMovie(&movie); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, movie)
}

func (h *MediaHandler) DeleteMovie(c *gin.Context) {
    id := c.Param("id")
    if err := h.MediaService.DeleteMovie(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}

// TVShow handlers
func (h *MediaHandler) GetTVShows(c *gin.Context) {
    tvShows, err := h.MediaService.GetTVShows()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tvShows)
}

func (h *MediaHandler) GetTVShow(c *gin.Context) {
    id := c.Param("id")
    tvShow, err := h.MediaService.GetTVShowByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "TV Show not found"})
        return
    }
    c.JSON(http.StatusOK, tvShow)
}

func (h *MediaHandler) CreateTVShow(c *gin.Context) {
    var tvShow models.TVShow
    if err := c.ShouldBindJSON(&tvShow); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Captura o usuário autenticado
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
        return
    }

    // Define o ID do usuário
    tvShow.UserID = userID.(uuid.UUID)

    if err := h.MediaService.CreateTVShow(&tvShow); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, tvShow)
}

func (h *MediaHandler) UpdateTVShow(c *gin.Context) {
    id := c.Param("id")
    var tvShow models.TVShow
    data, err := h.MediaService.GetTVShowByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "TV Show not found"})
        return
    }

    // Faça a assertiva de tipo
    tvShow, ok := data.(models.TVShow)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid TV show data"})
        return
    }

    if err := c.ShouldBindJSON(&tvShow); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    tvShow.ID = uuid.MustParse(id) // Garantir que o ID do programa de TV seja mantido
    if err := h.MediaService.UpdateTVShow(&tvShow); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tvShow)
}

func (h *MediaHandler) DeleteTVShow(c *gin.Context) {
    id := c.Param("id")
    if err := h.MediaService.DeleteTVShow(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}

// Add the handlers for fetching data from TheMovieDB
func (h *MediaHandler) GetTrendingAllWeekBR(c *gin.Context) {
    data, err := h.MediaService.GetTrendingAllWeekBR()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, data)
}

func (h *MediaHandler) GetTrendingAllDayBR(c *gin.Context) {
    data, err := h.MediaService.GetTrendingAllDayBR()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, data)
}

func (h *MediaHandler) SearchContent(c *gin.Context) {
    content := c.Param("content")
    data, err := h.MediaService.SearchContent(content)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, data)
}
