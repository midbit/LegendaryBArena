package service

import (
	interface_repository "LegendaryBArena/src/interface/repository"
	"LegendaryBArena/src/models"
	"LegendaryBArena/src/utility"

	"github.com/bwmarrin/discordgo"
)

type CardService struct {
	CardRepository interface_repository.CardRepositoryInterface
}

func (service CardService) FindOwnerCardByName(name string, user *discordgo.User) (*models.Card, error) {
	cards, err := service.CardRepository.FindOwnerCardByName(name, user.ID)
	if err != nil {
		return nil, err
	}
	if len(cards) == 0 {
		return nil, utility.CardDoesNotExistError{Card: name}
	}
	return &cards[0], nil
}
