package databaseOperations

import (
	"database/sql"
	"log"
)

func GetMessages(fromUserId int, toUserId int, db *sql.DB) []string {
	rows, err := db.Query("SELECT message FROM messages WHERE userFrom = ? AND userTo = ?", fromUserId, toUserId)

	var list []string

	for rows.Next() {
		var text string
		err = rows.Scan(&text)
		if err != nil {
			panic(err.Error())
		}
		list = append(list, text)
	}

	return list
}

func SendMessage(messageString string, fromUserId int, toUserId int, db *sql.DB) {
	_, err := db.Exec("INSERT INTO messages (message,userFrom,userTo,timeStamp) VALUES(?,?,?,NOW())", messageString, fromUserId, toUserId)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
