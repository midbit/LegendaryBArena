package channel

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func PrivateTextMessage(session *discordgo.Session, user *discordgo.User, message string) {
	channel, err := session.UserChannelCreate(user.ID)
	if err != nil {
		log.Print(err.Error())
	}
	session.ChannelMessageSend(channel.ID, message)

}

func PrivateSpecificMessage(session *discordgo.Session, user *discordgo.User, message *discordgo.MessageSend) {
	channel, err := session.UserChannelCreate(user.ID)
	if err != nil {
		log.Print(err.Error())
	}
	session.ChannelMessageSendComplex(channel.ID, message)
}

func SpecificMessage(session *discordgo.Session, channel_id string, message *discordgo.MessageSend) {
	session.ChannelMessageSendComplex(channel_id, message)
}
