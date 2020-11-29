package main

import (
	"LegendaryBArena/src/models"
	"LegendaryBArena/src/repository"
	"LegendaryBArena/src/server"
	"LegendaryBArena/src/service"
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

func createBot(db *gorm.DB) *server.Bot {
	userRepo := repository.UserRepository{
		Connection: db,
	}
	cardRepo := repository.CardRepository{
		Connection: db,
	}
	boosterRepo := repository.BoosterRepository{
		Connection: db,
	}
	userService := service.UserService{
		UserRepository: userRepo,
	}
	cardService := service.CardService{
		CardRepository: cardRepo,
	}
	boosterService := service.BoosterService{
		BoosterRepository: boosterRepo,
	}
	bot := server.Bot{
		CardService:    cardService,
		UserService:    userService,
		BoosterService: boosterService,
	}
	return &bot

}
func createSession(db *gorm.DB) *discordgo.Session {
	token := utility.GetToken()
	bot := createBot(db)
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
	}
	discord.AddHandler(bot.HandleMessage)

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
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Card{})
	return db
}

func main() {
	fmt.Println("Hello")
	loadEnvironment()
	db := connectDatabase()
	discord := createSession(db)
	err := discord.Open()
	if err != nil {
		log.Fatalf("Error Connection to Discord Server: %s", err.Error())
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	discord.Close()

}
