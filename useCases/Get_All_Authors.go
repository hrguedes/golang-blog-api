package useCases

import (
	"encoding/json"
	"log"
	"net/http"
	"web-service-gin/persistence"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	log.Println("Get All Posts")
	var authors = persistence.ListAllAuthors()
	// return
	json.NewEncoder(w).Encode(authors)
}
