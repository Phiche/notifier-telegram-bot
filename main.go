package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"time"
)

var clientToken = os.Getenv("CLIENT_TOKEN")

var clientBot, clientErr = tgbotapi.NewBotAPI(clientToken)

func main() {
	if clientErr != nil {
		log.Panic(clientErr)
	}

	clientBot.Debug = true

	log.Printf("Authorized on account %s", clientBot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := clientBot.GetUpdatesChan(u)

	if err != nil {
		log.Panic(clientErr)
	}

	adminMenu()

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		message := update.Message
		userFrom := message.From
		chatFrom := message.Chat

		if message.IsCommand() {
			switch command := message.Command(); command {
			case "start":
				user := User{userFrom.ID,
					chatFrom.ID,
					chatFrom.FirstName,
					chatFrom.LastName,
					userFrom.UserName,
					time.Unix(int64(message.Date), 0)}
				err := saveUser(user)
				log.Println(err)
				log.Println("start!")
			default:
				log.Println("Extra command!")
			}
		} else {
			log.Println("not command!")
		}
	}

}

func sendMessageToAll(message string) {
	usersIds := getChatsId()
	for _, id := range usersIds {
		msg := tgbotapi.NewMessage(id, message)
		log.Println("try to send to channel " + string(id))
		clientBot.Send(msg)
	}
}
