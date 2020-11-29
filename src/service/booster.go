package service

import (
	interface_repository "LegendaryBArena/src/interface/repository"
	"LegendaryBArena/src/models"
)

type BoosterService struct {
	BoosterRepository interface_repository.BoosterRepositoryInterface
}

func (service BoosterService) FindBoosterByName(name string) (*models.Booster, error) {
	booster, err := service.BoosterRepository.FindBoosterByName(name)
	if err != nil {
		return booster, err
	}
	return booster, nil
}

func (service BoosterService) FindAllBoosters(page int) ([]models.Booster, error) {
	offset := 5 * page
	boosters, err := service.BoosterRepository.FindAllBoosters(offset)
	if err != nil {
		return boosters, err
	}
	return boosters, nil
}
