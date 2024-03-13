package repo

import (
	"gorm.io/gorm"
	"tours-service.xws.com/model"
)

type CheckpointRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *CheckpointRepository) CreateCheckpoint(checkpoint *model.Checkpoint) error {
	dbResult := repo.DatabaseConnection.Create(checkpoint)
	if dbResult.Error != nil {
		panic(dbResult.Error)
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
