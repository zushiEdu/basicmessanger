package databaseOperations

import (
	"database/sql"
	"errors"
	"log"
	"main/smallFunctions"
	"main/types"
	"strconv"
)

// UserExistsId checks if a user exists with a provided id
func UserExistsId(id int, db *sql.DB) bool {
	rows, err := db.Query("SELECT id FROM users WHERE id = ?", id)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	if rows.Next() {
		return true
	}

	return false
}

// UserExistsEmail checks if a user exists with a provided email
func UserExistsEmail(email string, db *sql.DB) bool {
	rows, err := db.Query("SELECT email FROM users WHERE email = ?", email)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	if rows.Next() {
		return true
	}

	return false
}

func CreateUser(user types.User, db *sql.DB) (int, error) {
	rows, _ := db.Query("SELECT email from users WHERE email = ?", user.Email)
	if !rows.Next() {
		_, err := db.Exec("INSERT INTO users (first_name, last_name, email, password) VALUES(?,?,?,?)", user.FirstName, user.LastName, user.Email, user.Password)

		if err != nil {
			log.Fatalf("Error: %s", err)
		}

		rows, err = db.Query("SELECT id FROM users WHERE email = ?", user.Email)

		if err != nil {
			log.Fatalf("Error: %s", err)
		}

		var id int
		rows.Next()
		err = rows.Scan(&id)

		if err != nil {
			log.Fatalf("Error: %s", err)
		}

		return id, nil
	} else {
		// User with this email already exists
		return -1, errors.New("User with this email already exists")
	}
}

func GetIdList(tokenSignature string, db *sql.DB) []types.SmallUser {
	var toUserId string
	row := db.QueryRow("SELECT id from tokens WHERE signature = ?", tokenSignature)
	row.Scan(&toUserId)
	convertedId, _ := strconv.Atoi(toUserId)

	var idList []int
	rows, _ := db.Query("SELECT DISTINCT userFrom from messages where userTo = ?", convertedId)
	for rows.Next() {
		var tempId int
		rows.Scan(&tempId)
		idList = append(idList, tempId)
	}

	rows, _ = db.Query("SELECT DISTINCT userTo from messages where userFrom = ?", convertedId)
	for rows.Next() {
		var tempId int
		rows.Scan(&tempId)
		if !smallFunctions.Contains(tempId, idList) {
			idList = append(idList, tempId)
		}
	}

	var smallUsers []types.SmallUser
	for i := 0; i < len(idList); i++ {
		var newUser types.SmallUser
		row := db.QueryRow("SELECT first_name, last_name FROM users WHERE id = ?", idList[i])
		row.Scan(&newUser.FirstName, &newUser.LastName)
		newUser.Id = idList[i]
		smallUsers = append(smallUsers, newUser)
	}
	return smallUsers
}

func GetUser(email string, db *sql.DB) ([]types.User, error) {
	if email == "" {
		rows, _ := db.Query("SELECT * FROM users")

		var users []types.User

		for rows.Next() {
			var id, firstName, lastName, password string
			err := rows.Scan(&id, &firstName, &lastName, &email, &password)
			if err != nil {
				return []types.User{}, err
			}
			cleanedId, _ := strconv.Atoi(id)
			users = append(users, types.User{Id: cleanedId, FirstName: firstName, LastName: lastName, Email: email, Password: password})
		}

		return users, nil
	} else {
		row, _ := db.Query("SELECT * FROM users WHERE email = ?", email)
		if row.Next() {
			var id, firstName, lastName, password string
			err := row.Scan(&id, &firstName, &lastName, &email, &password)
			if err != nil {
				return []types.User{}, err
			}
			cleanedId, _ := strconv.Atoi(id)
			return []types.User{{Id: cleanedId, FirstName: firstName, LastName: lastName, Email: email, Password: password}}, nil
		} else {
			return []types.User{}, errors.New("User not found")
		}
	}
}

// EditUser Edits users in db argument based on ID provided in user argument
func EditUser(user types.User, db *sql.DB) error {
	// User with the new email already exists and is not itself
	rows, err := db.Query("SELECT email from users WHERE email = ? AND id != ?", user.Email, user.Id)

	if err != nil {
		log.Fatalf("Error: %s", err)
		return err
	}

	if rows.Next() != true {
		// User does not change its email to a conflicting one or is not changing it
		if UserExistsId(user.Id, db) {
			_, err = db.Exec("UPDATE users SET first_name = ?, last_name = ?, email = ?, password = ? WHERE id = ?", user.FirstName, user.LastName, user.Email, user.Password, user.Id)
			if err != nil {
				log.Fatalf("Error: %s", err)
				return err
			}
		}
	} else {
		return errors.New("user with this email already exists")
	}
	return nil
}

// DeleteUser Deletes user in db argument based on provided email argument
func DeleteUser(email string, db *sql.DB) error {
	if UserExistsEmail(email, db) {
		_, err := db.Exec("DELETE FROM users WHERE email = ?", email)
		if err != nil {
			log.Fatalf("Error: %s", err)
			return err
		}
	} else {
		return errors.New("user does not exist")
	}
	return nil
}
