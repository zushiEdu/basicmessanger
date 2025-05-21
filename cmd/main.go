package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"main/startup"
)

const dbUser = "dnuser"
const dbPass = "dbpass"
const dbSource = "10.0.0.2:3306"
const dbName = "dbname"

var db *sql.DB

func main() {
	startup.Server()

	//databaseOperations.CreateUser("John", "Doe", "jdoe3@domain.com", "password123", db)
	//databaseOperations.CreateUser("Jane", "Doe", "jdoe4@domain.com", "password123", db)
	//databaseOperations.SendMessage("Hello World!", 17, 18, db)
	//databaseOperations.GetMessages(17, 18, db)
	//databaseOperations.EditUser(types.User{Id: 17, FirstName: "Ethan", LastName: "Huber", Email: "etho8325@gmail.com", Password: "p123"}, db)
	//databaseOperations.DeleteUser(14, db)

}
