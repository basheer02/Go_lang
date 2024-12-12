package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)


func signup(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user) //it store the recieved data to the parameter, if some data missing it set as null, required field is must else raise an error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created", "user": user})
}

func login(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successfull", "token": token})
}

