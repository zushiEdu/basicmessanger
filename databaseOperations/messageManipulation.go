package databaseOperations

import (
	"database/sql"
	"errors"
	"log"
	"main/types"
)

func GetMessages(messageRequest types.MessageRequest, db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT message FROM messages WHERE userTo = ?", messageRequest.ToUser)

	var list []string

	for rows.Next() {
		var message string
		err = rows.Scan(&message)

		if err != nil {
			log.Fatalf("Error %s", err)
		}
		list = append(list, message)
	}

	return list, nil
}

func SendMessage(message types.Message, db *sql.DB) error {
	if UserExistsId(message.FromUser, db) && UserExistsId(message.ToUser, db) {
		_, err := db.Exec("INSERT INTO messages (message,userFrom,userTo,timeStamp) VALUES(?,?,?,NOW())", message.Message, message.FromUser, message.ToUser)
		if err != nil {
			log.Fatalf("Error: %s", err)
			return err
		} else {
			return nil
		}
	}
	return errors.New("could not send message")
}
