package repository

import (
	"LegendaryBArena/src/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	connection *gorm.DB
}

func (repository UserRepository) FindUserByName(username string) (*models.User, error) {
	var user *models.User
	result := repository.connection.Joins("Card").Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (repository UserRepository) CreateUser(user_id string, user_name string) error {
	user := models.User{Username: user_id, Id: user_id}
	result := repository.connection.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository UserRepository) AddCardToUser(user_id string, card_id int) error {
	var user *models.User
	var card *models.Card
	result := repository.connection.Where("Id = ?", user_id).First(&user)
	if result.Error != nil {
		return result.Error
	}
	result = repository.connection.Where("Id = ?", card_id).First(&card)
	if result.Error != nil {
		return result.Error
	}
	repository.connection.Model(&user).Association("Card").Append(&card)
	result = repository.connection.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository UserRepository) RemoveCardFromUser(user_id string, card_id int) error {
	var user *models.User
	var card *models.Card
	result := repository.connection.Where("Id = ?", user_id).First(&user)
	if result.Error != nil {
		return result.Error
	}
	result = repository.connection.Where("Id = ?", card_id).First(&card)
	if result.Error != nil {
		return result.Error
	}
	repository.connection.Model(&user).Association("Card").Delete(&card)
	result = repository.connection.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
