package routes

import (
	"github.com/gorilla/mux"
	"web-service-gin/useCases"
)

func PostRoute(router *mux.Router) {
	router.HandleFunc("/api/post/all", useCases.GetAllPosts).Methods("GET")
	router.HandleFunc("/api/post/add", useCases.CreateNewPost).Methods("POST")
}
