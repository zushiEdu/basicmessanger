package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/databaseOperations"
	"main/types"
	"net/http"
)

func CreateUserHandler(ctx *gin.Context) {
	fmt.Println("Got create user request")

	var user types.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		fmt.Println("Could not process request body")
		return
	}

	id, err := databaseOperations.CreateUser(user, databaseOperations.GetDB())
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success", "id": id})
	} else {
		// **Should** always be user with email already exists
		ctx.JSON(http.StatusConflict, gin.H{"message": err.Error()})
	}
}

func EditUserHandler(ctx *gin.Context) {
	fmt.Println("Got create edit request")

	var user types.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		fmt.Println("Could not process request body")
		return
	}

	err := databaseOperations.EditUser(user, databaseOperations.GetDB())
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success", "data": nil})
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"message": err.Error(), "data": nil})
	}
}
