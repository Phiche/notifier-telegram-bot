package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var host = os.Getenv("HOST")
var port = os.Getenv("PORT")
var user = os.Getenv("USER")
var password = os.Getenv("PASSWORD")
var dbname = os.Getenv("DBNAME")
var sslmode = os.Getenv("SSLMODE")

var dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

//todo: store history of messages and save whole telegram user extra fields as json for future updates
func saveUser(user User) error {

	//Подключаемся к БД
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	//Создаем SQL запрос
	sqlStatement := `INSERT INTO users(user_id, chat_id, first_name, last_name, username, start_date) VALUES($1, $2, $3, $4, $5, $6);`

	//Выполняем наш SQL запрос
	if _, err = db.Exec(sqlStatement, user.id, user.chatId, user.firstName, user.lastName, user.username, user.startDate); err != nil {
		return err
	}

	return nil
}

func getChatsId() []int64 {

	//Подключаемся к БД
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Создаем SQL запрос
	sqlStatement := `select chat_id from users;`
	//Выполняем наш SQL запрос
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var chatIds []int64
	for rows.Next() {
		var chatId int64
		err := rows.Scan(&chatId)
		if err != nil {
			log.Fatal(err)
		}

		chatIds = append(chatIds, chatId)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return chatIds
}
