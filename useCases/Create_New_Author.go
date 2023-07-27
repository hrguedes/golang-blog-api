package useCases

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"web-service-gin/entity"
	"web-service-gin/persistence"
)

func CreateNewAuthor(w http.ResponseWriter, r *http.Request) {
	log.Println("Insert new Author")
	var editor entity.Editor
	json.NewDecoder(r.Body).Decode(&editor)

	// valid
	if editor.Name == "" {
		http.Error(w, "Please, set name of Author", http.StatusBadRequest)
		return
	}

	// exists the author?
	var _, err = persistence.GetAuthorByName(editor.Name)
	if err != nil {
		// generate new id
		editor.Id = uuid.New()

		// insert new author
		var authors = persistence.AddNewAuthor(editor)
		json.NewEncoder(w).Encode(authors)
	} else {
		http.Error(w, "Author has exists", http.StatusBadRequest)
		return
	}
}
