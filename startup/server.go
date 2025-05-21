package startup

import "main/databaseOperations"

func Server() {
	db = databaseOperations.ConnectToDB(dbUser, dbPass, dbSource, dbName)
	databaseOperations.CloseDB(db)
}
