package routers

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main/api"
	"main/config"
	"time"
)

func Activate() {
	fmt.Println("Starting routers")
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*", "http://localhost:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/login", api.LoginHandler)

	router.POST("/users", api.CreateUserHandler)
	router.PUT("/users", api.EditUserHandler)
	router.GET("/users", api.GetUserHandler)
	router.DELETE("/users", api.DeleteUserHandler)

	router.POST("/message", api.CreateMessageHandler)
	router.GET("/message", api.GetMessageHandler)

	env := config.LoadEnv()
	err := router.Run(env.APIHost + ":" + env.APIPort)
	if err != nil {
		return
	}
}
