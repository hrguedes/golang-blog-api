package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"web-service-gin/routes"
)

func main() {
	// config log
	log.SetPrefix("Blog API: ")
	log.SetFlags(0)

	// http server
	r := mux.NewRouter()

	// routes
	routes.AuthRoute(r)
	routes.PostRoute(r)

	// start server
	log.Println("Listen in http://localhost:80")
	err := http.ListenAndServe(":80", r)
	if err != nil {
		return
	}
}
