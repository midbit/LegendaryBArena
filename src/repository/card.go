package repository

import (
	"LegendaryBArena/src/models"

	"gorm.io/gorm"
)

type CardRepository struct {
	Connection *gorm.DB
}

func (repository CardRepository) FindCardByName(name string) (*models.Card, error) {
	var card models.Card
	result := repository.Connection.Where("name = ?", name).First(&card)
	if result.Error != nil {
		return &card, result.Error
	}
	return &card, nil
}

func (repository CardRepository) FindOwnerCardByName(name string, owner string) ([]models.Card, error) {
	var cards []models.Card
	var user models.User

	err := repository.Connection.Where("Id = ?", owner).Model(&user).Where("name = ?", name).Association("Card").Find(&cards)
	if err != nil {
		return cards, err
	}
	return cards, nil
}

func (repository CardRepository) FindCardByBooster(name string) ([]*models.Card, error) {
	var booster models.Booster
	result := repository.Connection.Joins("Card").First(&booster, "name = ?", name)
	if result.Error != nil {
		return booster.Cards, result.Error
	}
	return booster.Cards, nil
}
