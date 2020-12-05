package repository

import (
	"LegendaryBArena/src/models"

	"gorm.io/gorm"
)

type BoosterRepository struct {
	Connection *gorm.DB
}

func (repository BoosterRepository) FindBoosterByName(name string) (*models.Booster, error) {
	var booster models.Booster
	result := repository.Connection.Preload("Cards").Where("name = ?", name).First(&booster)
	if result.Error != nil {
		return &booster, result.Error
	}
	return &booster, nil
}

func (repository BoosterRepository) FindAllBoosters(offset int) ([]models.Booster, error) {
	var boosters []models.Booster
	result := repository.Connection.Limit(5).Offset(offset).Find(&boosters)
	if result.Error != nil {
		return boosters, result.Error
	}
	return boosters, nil
}
