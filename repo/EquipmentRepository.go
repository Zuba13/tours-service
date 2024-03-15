package repo

import (
	"gorm.io/gorm"
	"tours-service.xws.com/model"
)

type TourEquipmentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourEquipmentRepository) SaveTourEquipment(newEquipment []model.Equipment, tourID int32) error {
	var currentEquipment []int32
	repo.DatabaseConnection.Model(&model.TourEquipment{}).Where("tour_id = ?", tourID).Pluck("equipment_id", &currentEquipment)

	var newEquipmentIDs []int32
	for _, equipment := range newEquipment {
		found := false
		for _, currentEquipmentID := range currentEquipment {
			if equipment.Id == currentEquipmentID {
				found = true
				break
			}
		}
		if !found {
			newEquipmentIDs = append(newEquipmentIDs, equipment.Id)
		}
	}

	for _, newEquipmentID := range newEquipmentIDs {
		repo.DatabaseConnection.Create(&model.TourEquipment{TourId: tourID, EquipmentId: newEquipmentID})
	}

	return nil
}

func (repo *TourEquipmentRepository) GetEquipment() []model.Equipment {
	var equipment []model.Equipment
	repo.DatabaseConnection.Model(&model.Equipment{}).Find(&equipment)
	return equipment
}
