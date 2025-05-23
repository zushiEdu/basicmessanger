package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/databaseOperations"
	"net/http"
)

func CreateUserHandler(ctx *gin.Context) {
	firstName := ctx.Query("firstName")
	lastName := ctx.Query("lastName")
	email := ctx.Query("email")
	password := ctx.Query("password")

	fmt.Println("Got create user request")

	_, err := databaseOperations.CreateUser(firstName, lastName, email, password, databaseOperations.GetDB())
	if err == nil {
		return
	} else {
		// Should always be user with email already exists
		ctx.JSON(http.StatusConflict, gin.H{"message": err.Error(), "data": nil})
	}
}
