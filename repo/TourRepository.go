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
