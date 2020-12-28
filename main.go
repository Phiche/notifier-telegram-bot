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
	Poller: spamProtected,
}

var clientBot, clientErr = tb.NewBot(prefClient)

func main() {
	go adminMenu()
	if clientErr != nil {
		log.Panic(clientErr)
	}

	log.Println("main job starts!")

	clientBot.Handle("/start", func(m *tb.Message) {
		log.Println("start handled!")
		sender := m.Sender
		chat := m.Chat

		user := User{sender.ID,
			chat.ID,
			sender.FirstName,
			sender.LastName,
			sender.Username,
			time.Unix(m.Unixtime, 0)}
		err := saveUser(user)
		log.Println(err)
	})
	log.Println("clientBot before start!")
	clientBot.Start()
	log.Println("clientBot started!")
}

func sendMessageToAll(message string) {
	usersIds := getChatsId()
	for _, id := range usersIds {
		clientBot.Send(tb.ChatID(id), message)
	}
}
