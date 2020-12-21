package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
	"strings"
)

var serverPort = os.Getenv("PUBLIC_PORT")
var publicURL = os.Getenv("PUBLIC_URL")

var webhook = &tb.Webhook{
	Listen:   ":" + serverPort,
	Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
}

var spamProtected = tb.NewMiddlewarePoller(webhook, func(upd *tb.Update) bool {
	if upd.Message == nil {
		return true
	}

	if strings.Contains(upd.Message.Text, "spam") {
		return false
	}

	return true
})
