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

func (service *TourService) Update(tour *model.Tour) error {
	err := service.TourRepo.UpdateTour(tour)
	if err != nil {
		fmt.Println("Error updating tour: ", err)
		return err
	}
	return nil
}

func (service *TourService) AddEquipment(tourId int32, newEquipment []model.Equipment) error {
	err := service.TourRepo.AddEquipment(tourId, newEquipment)
	if err != nil {
		fmt.Println("Error adding equipment to tour: ", err)
		return err
	}
	return nil
}
