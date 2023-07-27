package entity

import "github.com/google/uuid"

type Post struct {
	Id      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Author  Editor    `json:"author"`
	Content string    `json:"content"`
}
