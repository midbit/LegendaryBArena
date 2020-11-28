package interface_repository

import "LegendaryBArena/src/models"

type UserRepositoryInterface interface {
	FindUserByName(username string) (*models.User, error)
	CreateUser(user_id string, user_name string) error
	AddCardToUser(user_id string, card_id int) error
	RemoveCardFromUser(user_id string, card_id int) error
}
