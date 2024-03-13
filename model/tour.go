package model

import (
	"github.com/google/uuid"
)

type Difficult int

const (
	EASY Difficult = iota
	MEDIUM
	HARD
)

type Status int

const (
	DRAFT Status = iota
	PUBLISHED
	ARCHIVED
)

type Tour struct {
	ID          uuid.UUID `json:"id"`
	AuthorID    uuid.UUID `json:"author_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Difficult   Difficult `json:"difficult"`
}
