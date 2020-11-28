package repository

import "gorm.io/gorm"

type BoosterRepository struct {
	connection *gorm.DB
}
