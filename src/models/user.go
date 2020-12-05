package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       string  `gorm:"primaryKey"`
	Username string  `gorm:"index"`
	Card     []*Card `gorm:"many2many:user_card;"`
	Points   int     `gorm:"default:0;"`
}

func (user User) FindCards(name string) *Card {
	for _, card := range user.Card {
		if card.Name == name {
			return card
		}
	}
	return nil
}
