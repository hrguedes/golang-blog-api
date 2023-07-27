package persistence

import (
	"errors"
	"web-service-gin/entity"
)

var posts = []entity.Post{}

func ListAllPosts() []entity.Post {
	return posts
}

func AddNewPost(post entity.Post) []entity.Post {
	// insert new author
	posts = append(posts, post)
	return posts
}

func GetAuthorByTitle(title string) (entity.Post, error) {
	// exists the author?
	for i := range posts {
		if posts[i].Title == title {
			return posts[i], nil
		}
	}
	return entity.Post{}, errors.New("no Post found")
}
