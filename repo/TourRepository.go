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
