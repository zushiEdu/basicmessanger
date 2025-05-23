package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/api"
)

func Activate() {
	fmt.Println("Starting routers")
	router := gin.Default()

	router.POST("/users/", api.CreateUserHandler)
	router.PUT("/users/", api.EditUserHandler)

	router.Run("localhost:2345")
}
