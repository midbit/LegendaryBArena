package service

import (
	interface_repository "LegendaryBArena/src/interface/repository"

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
