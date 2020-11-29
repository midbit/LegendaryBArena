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
	userModel, err := service.UserRepository.FindUserByName(user.Username)
	if err != nil {
		return nil, err
	}
	return userModel, nil
}
