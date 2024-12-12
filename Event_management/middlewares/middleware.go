package middlewares

import (
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization") //get the authorization token

	if token == "" {
		// it abort the request, it is middleware so it doesn't directly cancel the request
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token) //verifying the token

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	context.Set("userId", userId)
	context.Next() // it executes the next requests
}