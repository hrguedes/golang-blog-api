package entity

import "github.com/google/uuid"

type Editor struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
