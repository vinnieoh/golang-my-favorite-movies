package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/middlewares"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/services"
)

func RegisterRoutes(router *gin.Engine, userService *services.UserService, mediaService *services.MediaService, secretKey string) {
	v1 := router.Group("/v1")
	{
		userHandler := NewUserHandler(userService)
		v1.GET("/users", userHandler.GetUsers)
		v1.GET("/users/:id", userHandler.GetUser)
		v1.POST("/users", userHandler.CreateUser)
		v1.PUT("/users/:id", userHandler.UpdateUser)
		v1.DELETE("/users/:id", userHandler.DeleteUser)

		v1.POST("/login", userHandler.Login)

		mediaHandler := NewMediaHandler(mediaService)
		
		// Rotas sem autenticação
		v1.GET("/movies", mediaHandler.GetMovies)
		v1.GET("/movies/:id", mediaHandler.GetMovie)
		v1.GET("/tvshows", mediaHandler.GetTVShows)
		v1.GET("/tvshows/:id", mediaHandler.GetTVShow)
		v1.GET("/trending-all-week-br", mediaHandler.GetTrendingAllWeekBR)
		v1.GET("/trending-all-day-br", mediaHandler.GetTrendingAllDayBR)
		v1.GET("/search-content/:content", mediaHandler.SearchContent)
		
		// Rotas com autenticação
		mediaRoutes := v1.Group("/")
		mediaRoutes.Use(middlewares.AuthMiddleware(secretKey))
		{
			mediaRoutes.POST("/movies", mediaHandler.CreateMovie)
			mediaRoutes.DELETE("/movies/:id", mediaHandler.DeleteMovie)

			mediaRoutes.POST("/tvshows", mediaHandler.CreateTVShow)
			mediaRoutes.DELETE("/tvshows/:id", mediaHandler.DeleteTVShow)
		}
	}
}
