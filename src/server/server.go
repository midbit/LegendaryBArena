package server

import (
	interface_service "LegendaryBArena/src/interface/service"
	"LegendaryBArena/src/utility"
	"log"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	UserService interface_service.UserServiceInterface
}

func (b Bot) HandleGuildCreated(session *discordgo.Session, guild *discordgo.GuildCreate) {
	members := guild.Members
	for _, member := range members {
		err := b.UserService.AddUser(member.User)
		if err != nil {
			session.ChannelMessageSend(guild.Channels[0].ID, err.Error())
			break
		}
	}
}
func (b Bot) HandleMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	command, err := utility.ParseMessage(message.Content)
	if err != nil {
		privateMessage(session, message.Author, err.Error())
	}

}

func privateMessage(session *discordgo.Session, user *discordgo.User, message string) {
	channel, err := session.UserChannelCreate(user.ID)
	if err != nil {
		log.Print(err.Error())
	}
	session.ChannelMessageSend(channel.ID, message)

}
