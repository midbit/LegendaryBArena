package interface_service

import (
	"github.com/bwmarrin/discordgo"
)

type UserServiceInterface interface {
	AddUser(*discordgo.User) error
}
