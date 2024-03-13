package service

import (
	"fmt"

	"tours-service.xws.com/model"
	"tours-service.xws.com/repo"
)

type TourService struct {
	TourRepo *repo.TourRepository
}

func (service *TourService) Create(tour *model.Tour) error {
	err := service.TourRepo.CreateTour(tour)
	if err != nil {
		fmt.Println("Error creating tour: ", err)
		return err
	}
	return nil
}

func (service *TourService) GetAuthorTours(authorId int32) []model.Tour {
	tours := service.TourRepo.GetAuthorTours(authorId)
	return tours
}
