package service

import (
	"tours-service.xws.com/model"
	"tours-service.xws.com/repo"
)

type TourExecutionService struct {
	TourExecutionRepo *repo.TourExecutionRepository
}

func (service *TourExecutionService) ExecuteTour(tourId int32, touristId int32) (model.TourExecution, error) {
	tourex, err := service.TourExecutionRepo.CreateTourExecution(tourId, touristId)
	if err != nil {
		return model.TourExecution{}, err
	}
	return tourex, nil
}
