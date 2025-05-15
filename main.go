package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

const dbUser = "dnuser"
const dbPass = "dbpass"
const dbSource = "10.0.0.2:3306"
const dbName = "dbname"

func main() {
	db := connectToDB()

	//sendMessage("Hello World!", 1, 2, db)
	//getMessages(1, 2, db)
	//id, err := createUser("Ethan", "Huber", "etho8325@gmail.com", "password123", db)

	db.Close()
}

func createUser(firstName string, lastName string, email string, password string, db *sql.DB) (int, error) {
	rows, _ := db.Query("SELECT email from users WHERE email = '" + email + "'")
	if rows.Next() != true {
		_, err := db.Exec("INSERT INTO users (first_name, last_name, email, password) VALUES('" + firstName + "','" + lastName + "','" + email + "','" + password + "')")
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		rows, err = db.Query("SELECT id FROM users WHERE email = '" + email + "'")

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

func editUser() {

}

func deleteUser(id int, db *sql.DB) {

}

func getMessages(fromUserId int, toUserId int, db *sql.DB) {
	rows, err := db.Query("SELECT message FROM messages WHERE userFrom = '" + strconv.Itoa(fromUserId) + "' AND userTo = '" + strconv.Itoa(toUserId) + "'")

	for rows.Next() {
		var text string
		err = rows.Scan(&text)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(text)
	}
}

func sendMessage(messageString string, fromUserId int, toUserId int, db *sql.DB) {
	_, err := db.Exec("INSERT INTO messages (message,userFrom,userTo,timeStamp) VALUES('" + messageString + "','" + strconv.Itoa(fromUserId) + "','" + strconv.Itoa(toUserId) + "',NOW())")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func connectToDB() *sql.DB {
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbSource+")/"+dbName)

	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot ping database because %s", err)
	}

	return db
}
