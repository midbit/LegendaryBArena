package service

import (
	interface_repository "LegendaryBArena/src/interface/repository"
	"LegendaryBArena/src/models"

	"github.com/bwmarrin/discordgo"
)

type UserService struct {
	UserRepository interface_repository.UserRepositoryInterface
}

func (service UserService) AddUser(user *discordgo.User) error {
	err := service.UserRepository.CreateUser(user.ID, user.Username)
	if err != nil {
		return err
	}
	return nil
}

func (service UserService) FindUser(user *discordgo.User) (*models.User, error) {
	targetUser, err := service.UserRepository.FindUserByName(user.Username)
	if err != nil {
		return targetUser, err
	}
	return targetUser, nil
}

func (service UserService) AddCard(user *discordgo.User, card *models.Card) error {
	err := service.UserRepository.AddCardToUser(user.Username, card.Id)
	return err
}

func (service UserService) AddCards(user *discordgo.User, cards []models.Card) error {
	var ids []uint
	for _, card := range cards {
		ids = append(ids, card.Id)
	}
	err := service.UserRepository.AddCardsToUser(user.ID, ids)
	return err
}
