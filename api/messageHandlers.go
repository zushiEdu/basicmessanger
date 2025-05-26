package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/databaseOperations"
	"main/types"
	"net/http"
)

func CreateMessageHandler(ctx *gin.Context) {
	fmt.Println("Got create message request")

	var message types.Message
	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		fmt.Println("Could not process request body")
		return
	}

	err := databaseOperations.SendMessage(message, databaseOperations.GetDB())
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"message": err.Error()})
	}
}

func GetMessageHandler(ctx *gin.Context) {
	fmt.Println("Got get user request")

	var email types.Email
	if err := ctx.ShouldBindJSON(&email); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		fmt.Println("Could not process request body")
		return
	}

	user, err := databaseOperations.GetUser(email.Email, databaseOperations.GetDB())
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success", "data": user})
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"message": err.Error()})
	}
}
