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
	password, id := databaseOperations.GetUserPass(request.Email, db)

	if password == request.Password {
		if databaseOperations.TokenExists(id, db) {
			token, err := databaseOperations.GetToken(id, db)
			if err == nil {
				ctx.JSON(http.StatusOK, gin.H{"message": "Success", "token": token.Signature})
			} else {
				ctx.JSON(http.StatusConflict, gin.H{"message": err.Error()})
			}
		} else {
			token, err := databaseOperations.MakeToken(id, db)
			if err == nil {
				ctx.JSON(http.StatusOK, gin.H{"message": "Success", "token": token})
			} else {
				ctx.JSON(http.StatusConflict, gin.H{"message": err.Error()})
			}
		}
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Incorrect Password"})
	}
}
