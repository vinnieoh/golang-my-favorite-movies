package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/vinnieoh/golang-my-favorite-movies/app/internal/services"
    "github.com/gorilla/mux"
)

func GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
    movies, err := services.GetAllMovies()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(movies)
}

func SetupRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/movies", GetMoviesHandler).Methods("GET")
    return r
}
