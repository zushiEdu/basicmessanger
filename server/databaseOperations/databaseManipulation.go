package databaseOperations

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func ConnectToDB(dbUser string, dbPass string, dbSource string, dbPort string, dbName string) *sql.DB {
	var err error
	fmt.Printf("Attempting to connect to database '%s'\n", dbName)
	db, err = sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbSource+":"+dbPort+")/"+dbName)

	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot ping database because %s", err)
	}

	return db
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() error {
	fmt.Println("Closing DB")
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}
