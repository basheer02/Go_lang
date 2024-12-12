package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest_api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvents(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Registration failed"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registration successfull"})

}

func cancelRegistration(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event"})
		return
	}

	fmt.Println(eventId, userId)

	err = event.Cancel(userId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Event cancellation failed"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event cancellation successfull"})

}
