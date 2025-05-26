package databaseOperations

import (
	"database/sql"
	"log"
	"main/types"
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

// TODO: implement a check so that the system does not try to send a message to/from a user that does not exist
func SendMessage(message types.Message, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO messages (message,userFrom,userTo,timeStamp) VALUES(?,?,?,NOW())", message.Message, message.FromUser, message.ToUser)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}
	return nil
}
