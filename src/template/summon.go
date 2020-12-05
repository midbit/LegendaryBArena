package template

import (
	"LegendaryBArena/src/models"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func CreateSummonTemplate(card *models.Card, user *models.User) *discordgo.MessageSend {
	message := discordgo.MessageSend{
		TTS: false,
		Embed: &discordgo.MessageEmbed{
			Title: fmt.Sprintf("%s have summon %s !", user.Username, card.Name),
			Type:  discordgo.EmbedTypeRich,
			Image: &discordgo.MessageEmbedImage{
				URL: card.Photo,
			},
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Name",
					Value:  card.Name,
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "Description",
					Value:  card.Description,
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "Class",
					Value:  card.Class,
					Inline: true,
				},
			},
		},
	}
	return &message
}
