package models

import (
	wr "github.com/mroth/weightedrand"
	"gorm.io/gorm"
)

type Booster struct {
	gorm.Model
	Name        string `gorm:"index"`
	Description string
	Cards       []*Card `gorm:"many2many:booster_card;"`
	Picture     string
	Id          uint `gorm:"primaryKey"`
}

type cardProbability struct {
	Card        *Card
	Probability float64
}

func (booster Booster) GenerateCard(times int) []Card {
	choice := booster.generateWeightChoice()
	chooser, _ := wr.NewChooser(choice...)
	cards := []Card{}
	if len(choice) != 0 {
		for time := 0; time < times; time++ {
			cards = append(cards, chooser.Pick().(Card))
		}
	}
	return cards
}

func (booster Booster) generateWeightChoice() []wr.Choice {
	choices := []wr.Choice{}
	for _, card := range booster.Cards {
		choices = append(choices, wr.Choice{Item: *card, Weight: card.Rarity})
	}
	return choices

}
