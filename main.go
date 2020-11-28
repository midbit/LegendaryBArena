package main

import (
	"LegendaryBArena/src/server"
	"LegendaryBArena/src/utility"

	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func loadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func createSession() *discordgo.Session {
	token := utility.GetToken()
	bot := server.Bot{}
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
	}
	discord.AddHandler(bot.HandleMessage)
	discord.AddHandler(bot.HandleGuildCreated)

	return discord
}

func connectDatabase() *gorm.DB {
	user := utility.GetMySQLUser()
	password := utility.GetMySQLPassword()
	host := utility.GetMySQLHost()
	port := utility.GetMySQLPort()
	database := utility.GetMySQLDatabase()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error Connection to Database: %s", err.Error())
	}
	return db
}

func main() {
	fmt.Println("Hello")
	loadEnvironment()
	connectDatabase()
	discord := createSession()
	err := discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	discord.Close()

}
