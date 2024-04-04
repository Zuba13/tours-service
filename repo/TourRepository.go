package repo

import (
	"gorm.io/gorm"
	"tours-service.xws.com/model"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourRepository) CreateTour(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Create(tour)
	if dbResult.Error != nil {
		panic(dbResult.Error)
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourRepository) GetAuthorTours(authorId int32) []model.Tour {
	var tours []model.Tour
	repo.DatabaseConnection.Preload("TourEquipment").Where("author_id = ?", authorId).Find(&tours)
	return tours
}

func (tourRepo *TourRepository) GetSuggestions(page int, pageSize int, checkpoints []int64) ([]model.Tour, error) {
	var resultTours []model.Tour

	var tours []model.Tour
	if err := tourRepo.DatabaseConnection.Where("status = ?", model.PUBLISHED).Preload("Checkpoints").Find(&tours).Error; err != nil {
		return nil, err
	}

	for _, tour := range tours {
		var tourCheckpointIDs []int64
		for _, checkpoint := range tour.Checkpoints {
			tourCheckpointIDs = append(tourCheckpointIDs, int64(checkpoint.Id))
		}

		allPresent := true
		for _, checkpointID := range checkpoints {
			found := false
			for _, tourCheckpointID := range tourCheckpointIDs {
				if checkpointID == tourCheckpointID {
					found = true
					break
				}
			}
			if !found {
				allPresent = false
				break
			}
		}

		if allPresent {
			resultTours = append(resultTours, tour)
		}
	}
	return resultTours, nil
}

func (repo *TourRepository) UpdateTour(tour *model.Tour) (*model.Tour, error) {
	tx := repo.DatabaseConnection.Begin()

	if err := tx.Save(tour).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var foundTour model.Tour
	if err := tx.Preload("TourEquipment").Find(&foundTour).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var existingEquipmentIds []int32
	for _, equipment := range foundTour.TourEquipment {
		existingEquipmentIds = append(existingEquipmentIds, equipment.Id)
	}

	var newEquipmentIds []int32
	for _, equipment := range tour.TourEquipment {
		newEquipmentIds = append(newEquipmentIds, equipment.Id)
	}

	var removedEquipmentIds []int32
	for _, existingEquipmentId := range existingEquipmentIds {
		found := false
		for _, newEquipmentId := range newEquipmentIds {
			if existingEquipmentId == newEquipmentId {
				found = true
				break
			}
		}
		if !found {
			removedEquipmentIds = append(removedEquipmentIds, existingEquipmentId)
		}
	}

	if len(removedEquipmentIds) > 0 {
		tx.Where("tour_id = ? AND equipment_id IN (?)", tour.Id, removedEquipmentIds).Delete(&model.TourEquipment{})
	}

	tx.Commit()

	return tour, nil
}

func (repo *TourRepository) AddEquipment(tourId int32, newEquipment []model.Equipment) error {

	var tour model.Tour
	if err := repo.DatabaseConnection.Where("id = ?", tourId).First(&tour).Error; err != nil {
		return err
	}
	tour.TourEquipment = append(tour.TourEquipment, newEquipment...)

	if err := repo.DatabaseConnection.Save(&tour).Error; err != nil {
		return err
	}

	return nil
}

func (repo *TourRepository) GetTourById(tourId int32) (*model.Tour, error) {
	var tour model.Tour
	if err := repo.DatabaseConnection.Where("id = ?", tourId).First(&tour).Error; err != nil {
		return nil, err
	}
	return &tour, nil
}
