package persistence

import (
	"errors"
	"github.com/google/uuid"
	"web-service-gin/entity"
)

var authors = []entity.Editor{}

func ListAllAuthors() []entity.Editor {
	return authors
}

func AddNewAuthor(author entity.Editor) []entity.Editor {
	// insert new author
	authors = append(authors, author)
	return authors
}

func GetAuthorById(id uuid.UUID) (entity.Editor, error) {
	// exists the author?
	for i := range authors {
		if authors[i].Id == id {
			return authors[i], nil
		}
	}
	return entity.Editor{}, errors.New("no Authors found")
}

func GetAuthorByName(name string) (entity.Editor, error) {
	// exists the author?
	for i := range authors {
		if authors[i].Name == name {
			return authors[i], nil
		}
	}
	return entity.Editor{}, errors.New("no Authors found")
}
