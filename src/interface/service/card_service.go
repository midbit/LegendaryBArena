package interface_service

import (
	"LegendaryBArena/src/models"

	"github.com/bwmarrin/discordgo"
)

type CardServiceInterface interface {
	FindOwnerCardByName(name string, user *discordgo.User) (*models.Card, error)
}
