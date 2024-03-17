package service

import (
	"fmt"

	"tours-service.xws.com/model"
	"tours-service.xws.com/repo"
)

type CheckpointService struct {
	CheckpointRepo *repo.CheckpointRepository
}

func (service *CheckpointService) Create(checkpoint *model.Checkpoint, tourId int32) error {
	err := service.CheckpointRepo.CreateCheckpoint(checkpoint, tourId)
	if err != nil {
		fmt.Println("Error creating checkpoint: ", err)
		return err
	}
	return nil
}

func (service *CheckpointService) GetCheckpoints(tourId int32) []model.Checkpoint {
	checkpoints, err := service.CheckpointRepo.GetCheckpoints(tourId)
	if err != nil {
		fmt.Println("Error getting checkpoints: ", err)
		return nil
	}
	return checkpoints
}
