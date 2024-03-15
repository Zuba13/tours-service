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
	repo.DatabaseConnection.Where("author_id = ?", authorId).Find(&tours)
	return tours
}

func (repo *TourRepository) UpdateTour(tour *model.Tour) (*model.Tour, error) {
	dbResult := repo.DatabaseConnection.Save(tour)
	if dbResult.Error != nil {
		panic(dbResult.Error)
	}
	println("Rows affected: ", dbResult.RowsAffected)
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
