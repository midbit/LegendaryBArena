package repository

import (
	"LegendaryBArena/src/models"

	"gorm.io/gorm"
)

type CardRepository struct {
	connection *gorm.DB
}

func (repository UserRepository) FindCardByName(name string) (*models.Card, error) {
	var card *models.Card
	result := repository.connection.Where("name = ?", name).First(&card)
	if result.Error != nil {
		return card, result.Error
	}
	return card, nil
}

func (repository UserRepository) FindCardByBooster(name string) ([]*models.Card, error) {
	var booster *models.Booster
	result := repository.connection.Joins("Card").First(&booster, "name = ?", name)
	if result.Error != nil {
		return booster.Cards, result.Error
	}
	return booster.Cards, nil
}
