package useCases

import (
	"encoding/json"
	"log"
	"net/http"
	"web-service-gin/persistence"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	log.Println("Get All Posts")
	var posts = persistence.ListAllPosts()
	// return
	json.NewEncoder(w).Encode(posts)
}
