package databaseOperations

import (
	"database/sql"
	"errors"
	"log"
	"main/config"
	"main/smallFunctions"
	"main/types"
	"time"
)

func TokenIsValid(signature string, db *sql.DB) bool {
	row := db.QueryRow("SELECT expiry FROM tokens WHERE signature=?", signature)

	if row == nil {
		return false
	}

	var expiry string
	err := row.Scan(&expiry)
	if err != nil {
		return false
	}

	expiryTime, _ := time.Parse("2006-01-02 15:04:05", expiry)
	if expiryTime.Before(time.Now()) {
		db.Exec("DELETE FROM tokens WHERE signature=?", signature)
		return false
	}

	return true
}

func TokenExists(id int, db *sql.DB) bool {
	_, err := GetToken(id, db)
	if err == nil {
		return true
	}
	return false
}

func GetToken(id int, db *sql.DB) (types.Token, error) {
	rows, err := db.Query("SELECT * FROM tokens WHERE id = ?", id)

	var signature, expiry string

	if rows != nil && rows.Next() {
		err = rows.Scan(&id, &signature, &expiry)
		if err != nil {
			log.Fatalf("Error %s", err)
		}
		return types.Token{ID: id, Signature: signature, Expiry: expiry}, nil
	} else {
		return types.Token{}, errors.New("Token not found")
	}
}

func MakeToken(id int, db *sql.DB) (string, error) {
	token := smallFunctions.GenerateToken()
	_, err := db.Exec("INSERT INTO tokens (id, signature, expiry) VALUES(?,?,DATE_ADD(NOW(), INTERVAL ? DAY))", id, token, config.ExpiryOffset)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return "", err
	} else {
		return token, nil
	}
}
