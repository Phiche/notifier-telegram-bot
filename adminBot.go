package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
	"strings"
	"time"
)

var adminToken = os.Getenv("ADMIN_TOKEN")

//var serverPort = os.Getenv("PORT")
//var publicURL = os.Getenv("PUBLIC_URL")

/*var webhook = &tb.Webhook{
Listen:   ":" + serverPort,
Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
}*/

var poller = &tb.LongPoller{Timeout: 15 * time.Second}
var spamProtected = tb.NewMiddlewarePoller(poller, func(upd *tb.Update) bool {
	if upd.Message == nil {
		return true
	}

	if strings.Contains(upd.Message.Text, "spam") {
		return false
	}

	return true
})

var pref = tb.Settings{
	Token: adminToken,
	//Poller: webhook,
	Poller: spamProtected,
}

var adminBot, adminErr = tb.NewBot(pref)

func adminMenu() {
	adminBot.Handle(tb.OnText, func(m *tb.Message) {
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

	adminBot.Start()
}
