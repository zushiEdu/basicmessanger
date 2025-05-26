package startup

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/config"
	"main/databaseOperations"
	"main/routers"
)

func Server() {
	env := config.LoadEnv()

	switch env.MODE {
	case "DEBUG":
		gin.SetMode(gin.DebugMode)
	case "PRODUCTION":
		gin.SetMode(gin.ReleaseMode)
	default:
		log.Fatalf("Invalid server mode is set")
	}

	SetAwaitTermination()
	databaseOperations.ConnectToDB(env.DBUser, env.DBPass, env.DBHost, env.DBPort, env.DBName)
	routers.Activate()
}
