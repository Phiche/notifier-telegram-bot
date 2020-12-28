package main

import (
	"os"
)

var serverPort = os.Getenv("PORT")
var databaseUrl = os.Getenv("DATABASE_URL")
var publicURL = os.Getenv("PUBLIC_URL")

/*var webhook = &tb.Webhook{
	Listen:   ":" + serverPort,
	Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL + "/botClient"},
}

var spamProtected = tb.NewMiddlewarePoller(webhook, func(upd *tb.Update) bool {
	if upd.Message == nil {
		return true
	}

	if strings.Contains(upd.Message.Text, "spam") {
		return false
	}

	return true
})*/
