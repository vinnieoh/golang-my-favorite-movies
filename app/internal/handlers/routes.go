package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vinnieoh/golang-my-favorite-movies/app/internal/services"
)

func RegisterRoutes(router *gin.Engine, userService *services.UserService) {
	v1 := router.Group("/v1")
	{
		userHandler := NewUserHandler(userService)
		v1.GET("/users", userHandler.GetUsers)
		v1.GET("/users/:id", userHandler.GetUser)
		v1.POST("/users", userHandler.CreateUser)
		v1.PUT("/users/:id", userHandler.UpdateUser)
		v1.DELETE("/users/:id", userHandler.DeleteUser)

		// movieHandler := NewMovieHandler(movieService)
		// v1.GET("/movies", movieHandler.GetMovies)
		// v1.GET("/movies/:id", movieHandler.GetMovie)
		// v1.POST("/movies", movieHandler.CreateMovie)
		// v1.PUT("/movies/:id", movieHandler.UpdateMovie)
		// v1.DELETE("/movies/:id", movieHandler.DeleteMovie)

		// tvShowHandler := NewTVShowHandler(tvShowService)
		// v1.GET("/tvshows", tvShowHandler.GetTVShows)
		// v1.GET("/tvshows/:id", tvShowHandler.GetTVShow)
		// v1.POST("/tvshows", tvShowHandler.CreateTVShow)
		// v1.PUT("/tvshows/:id", tvShowHandler.UpdateTVShow)
		// v1.DELETE("/tvshows/:id", tvShowHandler.DeleteTVShow)
	}
}
