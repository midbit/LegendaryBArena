package models

import "gorm.io/gorm"

type Booster struct {
	gorm.Model
	Name        string `gorm:"index"`
	Description string
	Cards       []*Card `gorm:"many2many:booster_card;"`
	Picture     string
	Id          uint `gorm:"primaryKey"`
}
