package router

import (
	"github.com/gorilla/mux"
	"github.com/karmsetu/maongo-api/controller"
)

func Router()*mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")

	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")

	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")

	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	
	router.HandleFunc("/api/delete/movies", controller.DeleteAllMovies).Methods("DELETE")
	// router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")


	

	return router
}