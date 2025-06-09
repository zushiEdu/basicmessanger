package databaseOperations

import (
	"database/sql"
	"errors"
	"log"
	"main/types"
)

func GetMessages(messageRequest types.MessageRequest, db *sql.DB) ([]types.MessageResponse, error) {
	rows, err := db.Query("SELECT message, userFrom FROM messages WHERE userTo = ? or userFrom = ?", messageRequest.InvolvingUser, messageRequest.InvolvingUser)

	var list []types.MessageResponse

	for rows.Next() {
		var message string
		var userFrom int
		err = rows.Scan(&message, &userFrom)

		if err != nil {
			log.Fatalf("Error %s", err)
		}
		list = append(list, types.MessageResponse{Message: message, UserFrom: userFrom})
	}

	return list, nil
}

func SendMessage(token string, message types.Message, db *sql.DB) error {
	row := db.QueryRow("SELECT id FROM tokens WHERE signature = ?", token)
	var fromId int
	row.Scan(&fromId)

	if UserExistsId(fromId, db) && UserExistsId(message.ToUser, db) {
		_, err := db.Exec("INSERT INTO messages (message,userFrom,userTo,timeStamp) VALUES(?,?,?,NOW())", message.Message, fromId, message.ToUser)
		if err != nil {
			log.Fatalf("Error: %s", err)
			return err
		} else {
			return nil
		}
	}
	return errors.New("could not send message")
}
