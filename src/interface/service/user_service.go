package interface_service

import (
	"LegendaryBArena/src/models"

	"github.com/bwmarrin/discordgo"
)

type UserServiceInterface interface {
	AddUser(*discordgo.User) error
	FindUser(*discordgo.User) (*models.User, error)
}
