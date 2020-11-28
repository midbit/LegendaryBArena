package models

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	Name        string `gorm:"index;unique"`
	Description string
	Photo       string
	Rarity      string
	Booster     []*Booster `gorm:"many2many:booster_card;"`
	User        []*User    `gorm:"many2many:user_card;"`
	Id          uint       `gorm:"primaryKey"`
}
