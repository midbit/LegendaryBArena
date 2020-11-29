package interface_repository

import "LegendaryBArena/src/models"

type BoosterRepositoryInterface interface {
	FindBoosterByName(name string) (*models.Booster, error)
	FindAllBoosters(offset int) ([]models.Booster, error)
}
