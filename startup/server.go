package startup

import (
	"github.com/gin-gonic/gin"
	"main/config"
	"main/databaseOperations"
	"main/routers"
)

func ProductionServer() {
	env := config.LoadEnv()
	gin.SetMode(gin.ReleaseMode)
	SetAwaitTermination()
	databaseOperations.ConnectToDB(env.DBUser, env.DBPass, env.DBHost, env.DBPort, env.DBName)
	routers.Activate()
}
