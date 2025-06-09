package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/databaseOperations"
	"main/types"
	"net/http"
)

func LoginHandler(ctx *gin.Context) {
	fmt.Println("Got login request")

	var request types.UserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		fmt.Println("Could not process request body")
		return
	}

	db := databaseOperations.GetDB()
	user, _ := databaseOperations.GetUser(request.Email, db)
	tokenRequest := types.TokenRequest{ID: user[0].Id}

	if user[0].Email == request.Email && user[0].Password == request.Password {
		if databaseOperations.TokenExists(tokenRequest, db) {
			ctx.JSON(http.StatusConflict, gin.H{"message": "Token already exists"})
		} else {
			token, err := databaseOperations.MakeToken(tokenRequest, db)
			if err == nil {
				ctx.JSON(http.StatusOK, gin.H{"message": "Success", "token": token})
			} else {
				ctx.JSON(http.StatusConflict, gin.H{"message": err.Error()})
			}
		}
	}
}
