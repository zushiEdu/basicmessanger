package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/databaseOperations"
	"main/types"
	"net/http"
	"strings"
)

func CreateUserHandler(ctx *gin.Context) {
	fmt.Println("Got create user request")

	var user types.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
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
	fmt.Println("Got edit user request")

	var user types.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		fmt.Println("Could not process request body")
		return
	}

	err := databaseOperations.EditUser(user, databaseOperations.GetDB())
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"message": err.Error()})
	}
}

func GetUserHandler(ctx *gin.Context) {
	fmt.Println("Got get user request")

	mode := ctx.Query("mode")
	email := ctx.Query("email")

	token := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")

	db := databaseOperations.GetDB()
	if mode == "single" {
		user, err := databaseOperations.GetUser(email, db)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{"message": "Success", "data": user})
		} else {
			ctx.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		}
	} else if mode == "multi" {
		list := databaseOperations.GetIdList(token, db)
		ctx.JSON(http.StatusOK, gin.H{"message": "Success", "data": list})
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"message": "Invalid request"})
	}
}

func DeleteUserHandler(ctx *gin.Context) {
	fmt.Println("Got delete user request")

	var email types.Email
	if err := ctx.ShouldBindJSON(&email); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		fmt.Println("Could not process request body")
		return
	}

	err := databaseOperations.DeleteUser(email.Email, databaseOperations.GetDB())
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"message": err.Error()})
	}
}
