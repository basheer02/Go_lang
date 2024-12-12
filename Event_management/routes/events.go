package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {

	userId := context.GetInt64("userId")

	events, err := models.GetAllEvents(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err, "err": err})
	}

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	userId := context.GetInt64("userId")

	event, err := models.GetEventByID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": " Not authorized to view the event"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event) //it store the recieved data to the parameter, if some data missing it set as null, required field is must else raise an error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	userId := context.GetInt64("userId")

	event.UserID = userId
	err = event.Save()

	if err != nil {
		fmt.Println(" Error in storing event")
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func updateEvent(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	userId := context.GetInt64("userId")

	res, err := models.GetEventByID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	if res.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": " Not authorized to update the event"})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse the data."})
		return
	}

	updatedEvent.ID = id

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event updated", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	userId := context.GetInt64("userId")
	res, err := models.GetEventByID(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	if res.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": " Not authorized to delete the event"})
		return
	}

	err = res.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted"})
}
