package service

import (
	"fmt"

	"tours-service.xws.com/model"
	"tours-service.xws.com/repo"
)

type CheckpointService struct {
	CheckpointRepo *repo.CheckpointRepository
}

func (service *CheckpointService) Create(checkpoint *model.Checkpoint) error {
	err := service.CheckpointRepo.CreateCheckpoint(checkpoint)
	if err != nil {
		fmt.Println("Error creating checkpoint: ", err)
		return err
	}
	return nil
}
