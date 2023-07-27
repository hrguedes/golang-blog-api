package routes

import (
	"github.com/gorilla/mux"
	"web-service-gin/useCases"
)

func AuthRoute(router *mux.Router) {
	router.HandleFunc("/api/author/all", useCases.GetAuthors).Methods("GET")
	router.HandleFunc("/api/author/add", useCases.CreateNewAuthor).Methods("POST")
}
