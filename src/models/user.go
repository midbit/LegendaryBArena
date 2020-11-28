package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       string  `gorm:"primaryKey"`
	Username string  `gorm:"index"`
	Card     []*Card `gorm:"many2many:user_card;"`
	Points   int     `gorm:"default:0;"`
}
