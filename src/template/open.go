package template

import (
	"LegendaryBArena/src/models"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func CreateOpenTemplate(cards []models.Card, booster models.Booster) *discordgo.MessageSend {
	var emmbededField []*discordgo.MessageEmbedField
	for _, card := range cards {
		emmbededField = append(emmbededField, &discordgo.MessageEmbedField{
			Name:   "Card",
			Value:  card.Name,
			Inline: true,
		})
	}
	message := discordgo.MessageSend{
		TTS: false,
		Embed: &discordgo.MessageEmbed{
			Title:  fmt.Sprintf("You have opend the %s booster", booster.Name),
			Type:   discordgo.EmbedTypeRich,
			Fields: emmbededField,
		},
	}
	return &message
}
