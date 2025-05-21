package databaseOperations

import (
	"database/sql"
	"errors"
	"log"
	"main/types"
)

func CreateUser(firstName string, lastName string, email string, password string, db *sql.DB) (int, error) {
	rows, _ := db.Query("SELECT email from users WHERE email = ?", email)
	if rows.Next() != true {
		_, err := db.Exec("INSERT INTO users (first_name, last_name, email, password) VALUES(?,?,?,?)", firstName, lastName, email, password)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		rows, err = db.Query("SELECT id FROM users WHERE email = ?", email)

		var id int
		rows.Next()
		err = rows.Scan(&id)
		if err != nil {
			panic(err.Error())
		}

		return id, nil
	} else {
		// User with this email already exists
		return -1, errors.New("user with this email already exists")
	}
}

func EditUser(user types.User, db *sql.DB) {
	_, err := db.Exec("UPDATE users SET first_name = ?, last_name = ?, email = ?, password = ? WHERE id = ?", user.FirstName, user.LastName, user.Email, user.Password, user.Id)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func DeleteUser(id int, db *sql.DB) {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
