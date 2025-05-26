package databaseOperations

import (
	"database/sql"
	"errors"
	"log"
	"main/types"
	"strconv"
)

func CreateUser(user types.User, db *sql.DB) (int, error) {
	rows, _ := db.Query("SELECT email from users WHERE email = ?", user.Email)
	if !rows.Next() {
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

func GetUser(email string, db *sql.DB) (types.User, error) {
	row, _ := db.Query("SELECT * FROM users WHERE email = ?", email)
	if row.Next() {
		var id, firstName, lastName, password string
		err := row.Scan(&id, &firstName, &lastName, &email, &password)
		if err != nil {
			return types.User{}, err
		}
		cleanedId, _ := strconv.Atoi(id)
		return types.User{Id: cleanedId, FirstName: firstName, LastName: lastName, Email: email, Password: password}, nil
	} else {
		return types.User{}, errors.New("no user exists with this email")
	}
}

// EditUser Edits users in db argument based on ID provided in user argument
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

// DeleteUser Deletes user in db argument based on provided email argument
func DeleteUser(email string, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE email = ?", email)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}
	return nil
}
