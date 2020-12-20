package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"time"
	//tb "gopkg.in/tucnak/telebot.v2"
)

var clientToken = os.Getenv("CLIENT_TOKEN")
var adminToken = os.Getenv("ADMIN_TOKEN")

var clientBot, clientErr = tgbotapi.NewBotAPI(clientToken)

func main() {
	//bot, err := tgbotapi.NewBotAPI("1414955792:AAGflb75_tjCI_o5ifCdpT9cmxOLoOsgMzs")
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

	adminBotJob()

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
			case "e", "export":
				log.Println("ne start!")
			default:
				log.Println("defaultt!")
			}
		} else {
			log.Println("not command!")
		}

		/*       log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)*/
	}
}

func adminBotJob() {
	bot, err := tgbotapi.NewBotAPI(adminToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		message := update.Message
		userFrom := message.From

		if message.IsCommand() {
			//check admin
			if userFrom.ID != 91469086 {
				log.Println("Wrong user!")
				log.Println(message)
				return
			}
			switch command := message.Command(); command {
			case "addnews":
				usersIds := getChatsId()
				for _, id := range usersIds {
					msg := tgbotapi.NewMessage(id, update.Message.Text)
					log.Println("try to send to channel " + string(id))
					clientBot.Send(msg)
				}
				log.Println(err)
			case "e", "export":
				log.Println("ne start!")
			default:
				log.Println("defaultt!")
			}
		} else {
			log.Println("not command!")
		}
	}
}
