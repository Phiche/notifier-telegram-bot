package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"strings"
)

//var poller = &tb.LongPoller{Timeout: 15 * time.Second}

func adminMenu() {
	var adminToken = os.Getenv("ADMIN_TOKEN")

	var adminWebhook = &tb.Webhook{
		Listen:   ":" + serverPort,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL + "/botAdmin"},
	}

	var prefAdmin = tb.Settings{
		Token:  adminToken,
		Poller: adminWebhook,
	}

	var adminBot, adminErr = tb.NewBot(prefAdmin)

	if adminErr != nil {
		log.Panic(adminErr)
	}

	log.Println("admin job starts!")
	adminBot.Handle(tb.OnText, func(m *tb.Message) {
		log.Println("admin handled on text!")
		if !m.Private() {
			return
		}
		if m.Chat.ID != 91469086 {
			return
		}
		if strings.HasPrefix(m.Text, "Постик:") {
			sendMessageToAll(strings.TrimPrefix(m.Text, "Постик:"))
		}
	})

	log.Println("adminBot before start!")
	adminBot.Start()
	log.Println("adminBot started!")
}
