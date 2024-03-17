package model

import "time"

type CheckpointStatus struct {
	TourExecutionId int32     `json:"TourExecutionId"`
	CheckpointId    int32     `json:"CheckpointId"`
	IsCompleted     bool      `json:"IsCompleted"`
	CompletionTime  time.Time `json:"CompletionTime"`
}
