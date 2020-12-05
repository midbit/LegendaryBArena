package server

import (
	"LegendaryBArena/src/channel"
	interface_service "LegendaryBArena/src/interface/service"
	"LegendaryBArena/src/template"
	"LegendaryBArena/src/utility"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	UserService    interface_service.UserServiceInterface
	CardService    interface_service.CardServiceInterface
	BoosterService interface_service.BoosterServiceInterface
}

func (b Bot) messageCheckDecorator(wrapped func(*discordgo.Session, *discordgo.MessageCreate), message *discordgo.MessageCreate) func(session *discordgo.Session, message *discordgo.MessageCreate) {
	_, err := b.UserService.FindUser(message.Author)
	if err != nil {
		b.UserService.AddUser(message.Author)
	}
	return wrapped

}

func (b Bot) handleMessageResult(session *discordgo.Session, message *discordgo.MessageCreate) {
	command, err := utility.ParseMessage(message.Content)
	if err != nil {
		channel.PrivateTextMessage(session, message.Author, err.Error())
		return
	}
	var messageSend *discordgo.MessageSend
	switch command.Command {
	case "help":
		messageSend = template.CreateHelpTemplate()
		channel.PrivateSpecificMessage(session, message.Author, messageSend)
		break
	case "summon":
		break
	case "open":
		name := strings.Join(command.Arguments[:], " ")
		booster, err := b.BoosterService.FindBoosterByName(name)
		if err != nil {
			channel.PrivateTextMessage(session, message.Author, err.Error())
			break
		}
		log.Print(booster.Cards)
		cards := booster.GenerateCard(5)
		err = b.UserService.AddCards(message.Author, cards)
		if err != nil {
			channel.PrivateTextMessage(session, message.Author, err.Error())
			break
		}
		messageSend = template.CreateOpenTemplate(cards, *booster)
		channel.PrivateSpecificMessage(session, message.Author, messageSend)

	case "boosters":
		boosters, err := b.BoosterService.FindAllBoosters(0)
		if err != nil {
			channel.PrivateTextMessage(session, message.Author, err.Error())
			break
		}
		messageSend = template.CreateBoosterTemplate(boosters)
		channel.PrivateSpecificMessage(session, message.Author, messageSend)
		break
	default:
		break
	}
}
func (b Bot) HandleMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	b.messageCheckDecorator(b.handleMessageResult, message)(session, message)
}
