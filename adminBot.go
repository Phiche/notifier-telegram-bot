package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"strings"
)

var adminToken = os.Getenv("ADMIN_TOKEN")

//var poller = &tb.LongPoller{Timeout: 15 * time.Second}

var prefAdmin = tb.Settings{
	Token:  adminToken,
	Poller: webhook,
	//Poller: spamProtected,
}

var adminBot, adminErr = tb.NewBot(prefAdmin)

func adminMenu() {
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
