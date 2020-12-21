package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

var clientToken = os.Getenv("CLIENT_TOKEN")

var prefClient = tb.Settings{
	Token:  clientToken,
	Poller: webhook,
	//Poller: spamProtected,
}

var clientBot, clientErr = tb.NewBot(prefClient)

func main() {
	if clientErr != nil {
		log.Panic(clientErr)
	}

	migrateDb()

	adminMenu()

	clientBot.Handle("/start", func(m *tb.Message) {
		userJoined := m.UserJoined
		chat := m.Chat

		user := User{userJoined.ID,
			chat.ID,
			userJoined.FirstName,
			userJoined.LastName,
			userJoined.Username,
			time.Unix(m.Unixtime, 0)}
		err := saveUser(user)
		log.Fatal(err)
	})

	clientBot.Start()
}

func sendMessageToAll(message string) {
	usersIds := getChatsId()
	for _, id := range usersIds {
		clientBot.Send(tb.ChatID(id), message)
	}
}

func migrateDb() {
	m, err := migrate.New(
		"file://migrations",
		databaseUrl+"?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
