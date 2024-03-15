package service

import (
	"tours-service.xws.com/model"
	"tours-service.xws.com/repo"
)

type EquipmentService struct {
	EquipmentRepo *repo.TourEquipmentRepository
}

func (service *EquipmentService) SaveTourEquipment(newEquipment []model.Equipment, tourId int32) error {
	err := service.EquipmentRepo.SaveTourEquipment(newEquipment, tourId)
	if err != nil {
		return err
	}
	return nil
}

func (service *EquipmentService) GetEquipment() []model.Equipment {
	equipment := service.EquipmentRepo.GetEquipment()
	return equipment
}
