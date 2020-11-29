package interface_repository

import "LegendaryBArena/src/models"

type CardRepositoryInterface interface {
	FindCardByName(name string) (*models.Card, error)
}
