package template

import (
	"github.com/bwmarrin/discordgo"
)

func CreateHelpTemplate() *discordgo.MessageSend {
	message := discordgo.MessageSend{
		TTS: false,
		Embed: &discordgo.MessageEmbed{
			Title: "List of available commands",
			Type:  discordgo.EmbedTypeRich,
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "help",
					Value:  "Generate a list of commands",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "open {Name}",
					Value:  "Open a booster of the specified name",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "boosters",
					Value:  "List all available boosters",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "summon {Name}",
					Value:  "Summon a hero on to the channel",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "exchange {Name} {UserName}",
					Value:  "Exchange a specified card with a specified user",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "me",
					Value:  "Show my stats",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "collections",
					Value:  "Show my collections",
					Inline: true,
				},
			},
		},
	}
	return &message
}
