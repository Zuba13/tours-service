package model

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
	Id          int32     `json:"id"`
	AuthorId    int32     `json:"author_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Difficult   Difficult `json:"difficult"`
	Status      Status    `json:"status"`
}
