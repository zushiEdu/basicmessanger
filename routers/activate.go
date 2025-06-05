package routers

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main/api"
	"time"
)

func Activate() {
	fmt.Println("Starting routers")
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/users/", api.CreateUserHandler)
	router.PUT("/users/", api.EditUserHandler)
	router.GET("/users/", api.GetUserHandler)
	router.DELETE("/users/", api.DeleteUserHandler)

	router.POST("/message/", api.CreateMessageHandler)
	router.GET("/message/", api.GetMessageHandler)

	err := router.Run("10.0.0.243:2345")
	if err != nil {
		return
	}
}
