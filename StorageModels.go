package main

import "time"

type User struct {
	id        int
	chatId    int64
	firstName string
	lastName  string
	username  string
	startDate time.Time
}
