package template

import (
	"LegendaryBArena/src/models"

	"github.com/bwmarrin/discordgo"
)

func CreateBoosterTemplate(boosters []models.Booster) *discordgo.MessageSend {
	var emmbededField []*discordgo.MessageEmbedField
	for _, booster := range boosters {
		emmbededField = append(emmbededField, &discordgo.MessageEmbedField{
			Name:   "Booster",
			Value:  booster.Name,
			Inline: true,
		})
	}
	message := discordgo.MessageSend{
		TTS: false,
		Embed: &discordgo.MessageEmbed{
			Title:  "List of available boosters",
			Type:   discordgo.EmbedTypeRich,
			Fields: emmbededField,
		},
	}
	return &message
}
