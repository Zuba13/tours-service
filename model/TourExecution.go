package model

import "time"

type TourExecutionStatus int

const (
	Active TourExecutionStatus = iota
	Completed
	Abandoned
)

type TourExecution struct {
	Id                 int32               `json:"Id"`
	TouristId          int32               `json:"TouristId"`
	TourId             int32               `json:"TourId"`
	IsComposite        bool                `json:"IsComposite"`
	LastActivity       time.Time           `json:"LastActivity"`
	Status             TourExecutionStatus `json:"Status"`
	CheckpointStatuses []CheckpointStatus  `json:"CheckpointStatuses"`
	CoveredDistance    float32             `json:"CoveredDistance"`
	Tour               *Tour               `json:"tour"`
}
