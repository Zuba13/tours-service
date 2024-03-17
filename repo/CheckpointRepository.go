package repo

import (
	"gorm.io/gorm"
	"tours-service.xws.com/model"
)

type CheckpointRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *CheckpointRepository) CreateCheckpoint(checkpoint *model.Checkpoint, tourId int32) error {
	checkpoint.TourId = tourId
	dbResult := repo.DatabaseConnection.Create(checkpoint)
	if dbResult.Error != nil {
		panic(dbResult.Error)
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *CheckpointRepository) GetCheckpoints(tourId int32) ([]model.Checkpoint, error) {
	var checkpoints []model.Checkpoint
	if err := repo.DatabaseConnection.Where("tour_id = ?", tourId).Find(&checkpoints).Error; err != nil {
		return nil, err
	}
	return checkpoints, nil
}
