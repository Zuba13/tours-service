package repo

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"tours-service.xws.com/model"
)

type TourExecutionRepository struct {
	DatabaseConnection *gorm.DB
	CheckpointRepo     *CheckpointRepository
}

func (repo *TourExecutionRepository) CreateTourExecution(tourId int32, touristId int32) (model.TourExecution, error) {

	checkpoints, err := repo.CheckpointRepo.GetCheckpoints(tourId)
	if err != nil {
		return model.TourExecution{}, err
	}

	println("Checkpoints: ", checkpoints)

	if len(checkpoints) < 2 {
		return model.TourExecution{}, errors.New("tura mora imati najmanje 2 checkpointa")
	}

	tourExecution := model.TourExecution{
		TourId:          tourId,
		TouristId:       touristId,
		IsComposite:     false,
		LastActivity:    time.Now(),
		Status:          model.Active,
		CoveredDistance: 0.0,
	}

	if err := repo.DatabaseConnection.FirstOrCreate(&tourExecution, model.TourExecution{
		TourId:    tourId,
		TouristId: touristId,
	}).Error; err != nil {
		return model.TourExecution{}, err
	}

	for _, checkpoint := range checkpoints {
		checkpointStatus := model.CheckpointStatus{
			TourExecutionId: tourExecution.Id,
			CheckpointId:    checkpoint.Id,
			IsCompleted:     false,
			CompletionTime:  time.Time{},
		}
		if err := repo.DatabaseConnection.Create(&checkpointStatus).Error; err != nil {
			return model.TourExecution{}, err
		}
	}

	println("Tour execution created successfully")
	return tourExecution, nil
}
