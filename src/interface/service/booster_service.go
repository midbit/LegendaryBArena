package interface_service

import (
	"LegendaryBArena/src/models"
)

type BoosterServiceInterface interface {
	FindAllBoosters(page int) ([]models.Booster, error)
	FindBoosterByName(name string) (*models.Booster, error)
}
