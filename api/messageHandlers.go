package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/databaseOperations"
	"main/types"
	"net/http"
	"strconv"
	"strings"
)

func CreateMessageHandler(ctx *gin.Context) {
	fmt.Println("Got create message request")

	var message types.Message
	if err := ctx.BindJSON(&message); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		fmt.Println("Could not process request body")
		return
	}

	token := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")

	db := databaseOperations.GetDB()
	if databaseOperations.TokenIsValid(token, db) {
		err := databaseOperations.SendMessage(token, message, db)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
		} else {
			ctx.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		}
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"message": "Token is invalid"})
	}
}

func GetMessageHandler(ctx *gin.Context) {
	fmt.Println("Got get message request")

	var message types.MessageRequest
	message.InvolvingUser, _ = strconv.Atoi(ctx.Query("involvingUser"))
	message.Token = strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")

	db := databaseOperations.GetDB()
	if databaseOperations.TokenIsValid(message.Token, db) {
		messages, err := databaseOperations.GetMessages(message, db)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{"message": "Success", "data": messages})
		} else {
			ctx.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		}
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"message": "Token is invalid"})
	}
}
