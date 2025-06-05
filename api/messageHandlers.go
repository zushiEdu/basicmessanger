package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/databaseOperations"
	"main/types"
	"net/http"
	"strconv"
)

func CreateMessageHandler(ctx *gin.Context) {
	fmt.Println("Got create message request")

	var message types.Message
	if err := ctx.BindJSON(&message); err != nil {
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
	fmt.Println("Got get message request")

	var message types.MessageRequest
	message.ToUser, _ = strconv.Atoi(ctx.Query("toUser"))

	messages, err := databaseOperations.GetMessages(message, databaseOperations.GetDB())
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success", "data": messages})
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"message": err.Error()})
	}
}
