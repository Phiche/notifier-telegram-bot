package main

import (
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

	log.Println("main job starts!")

	adminMenu()

	clientBot.Handle("/start", func(m *tb.Message) {
		log.Println("start handled!")
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
	log.Println("clientBot started!")
}

func sendMessageToAll(message string) {
	usersIds := getChatsId()
	for _, id := range usersIds {
		clientBot.Send(tb.ChatID(id), message)
	}
}
