package databaseOperations

import (
	"database/sql"
	"log"
)

func ConnectToDB(dbUser string, dbPass string, dbSource string, dbName string) *sql.DB {
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbSource+")/"+dbName)

	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot ping database because %s", err)
	}

	return db
}

func CloseDB(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}
