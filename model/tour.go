package model

import (
	"time"
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
	Id                  int32                 `json:"Id"`
	AuthorId            int32                 `json:"AuthorId"`
	Name                string                `json:"Name"`
	Description         string                `json:"Description"`
	Tags                string                `gorm:"type:text" json:"Tags"`
	Difficult           Difficult             `json:"Difficult"`
	Price               float64               `json:"Price"`
	Status              Status                `json:"Status"`
	PublishTime         time.Time             `json:"PublishTime"`
	ArchiveTime         time.Time             `json:"ArchiveTime"`
	Checkpoints         []Checkpoint          `json:"Checkpoints"`
	Distance            float64               `json:"Distance"`
	TravelTimeAndMethod []TravelTimeAndMethod `gorm:"type:jsonb" json:"TravelTimeAndMethod"`
	TourEquipment       []Equipment           `gorm:"many2many:tour_equipments;" json:"TourEquipment"`
}
