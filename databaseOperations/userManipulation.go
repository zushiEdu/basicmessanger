package databaseOperations

import (
	"database/sql"
	"errors"
	"log"
	"main/types"
)

func CreateUser(user types.User, db *sql.DB) (int, error) {
	rows, _ := db.Query("SELECT email from users WHERE email = ?", user.Email)
	if rows.Next() != true {
		_, err := db.Exec("INSERT INTO users (first_name, last_name, email, password) VALUES(?,?,?,?)", user.FirstName, user.LastName, user.Email, user.Password)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		rows, err = db.Query("SELECT id FROM users WHERE email = ?", user.Email)

		var id int
		rows.Next()
		err = rows.Scan(&id)
		if err != nil {
			panic(err.Error())
		}

		return id, nil
	} else {
		// User with this email already exists
		return -1, errors.New("User with this email already exists")
	}
}

func EditUser(user types.User, db *sql.DB) error {
	// User with the new email already exists and is not itself
	rows, err := db.Query("SELECT email from users WHERE email = ? AND id != ?", user.Email, user.Id)
	if rows.Next() != true {
		// User does not change its email to a conflicting one or is not changing it
		rows, err = db.Query("SELECT id FROM users WHERE id = ?", user.Id)

		if err != nil {
			log.Fatalf("Error: %s", err)
			return err
		}

		if rows.Next() {
			_, err = db.Exec("UPDATE users SET first_name = ?, last_name = ?, email = ?, password = ? WHERE id = ?", user.FirstName, user.LastName, user.Email, user.Password, user.Id)
			if err != nil {
				log.Fatalf("Error: %s", err)
				return err
			}
			return nil
		} else {
			return errors.New("user with this ID does not exist")
		}
	} else {
		return errors.New("user with this email already exists")
	}
}

func DeleteUser(id int, db *sql.DB) {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
