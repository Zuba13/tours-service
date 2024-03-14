package service

import (
	"tours-service.xws.com/model"
	"tours-service.xws.com/repo"
)

type EquipmentService struct {
	repo *repo.TourEquipmentRepository
}

func (service *EquipmentService) SaveTourEquipment(newEquipment []model.Equipment, tourId int32) error {
	err := service.repo.SaveTourEquipment(newEquipment, tourId)
	if err != nil {
		return err
	}
	return nil
}
