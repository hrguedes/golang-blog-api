package useCases

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"web-service-gin/entity"
	"web-service-gin/persistence"
)

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	log.Println("Insert new Post")

	// ready body and convert o Post
	var post entity.Post
	var err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		if err.Error() == "invalid UUID length: 0" {
			http.Error(w, "don`t send not use parameter", http.StatusBadRequest)
			return
		}
	}

	// valid
	if post.Author.Id.String() == "" {
		http.Error(w, "Please input the Id of post author", http.StatusBadRequest)
		return
	}
	if post.Content == "" {
		http.Error(w, "Please input the Content post", http.StatusBadRequest)
		return
	}
	if post.Title == "" {
		http.Error(w, "Please input the Title post", http.StatusBadRequest)
		return
	}

	// exists the author?
	var author, errAuthor = persistence.GetAuthorById(post.Author.Id)
	if errAuthor == nil {
		// generate new id
		post.Id = uuid.New()
		post.Author = author

		// insert new post
		var posts = persistence.AddNewPost(post)

		// return
		json.NewEncoder(w).Encode(posts)
	} else {
		http.Error(w, "Please input Id Author exists", http.StatusBadRequest)
		return
	}
}
