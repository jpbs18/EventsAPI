package routes

import (
	"Events-API/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event with that id"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event"})
		return
	}

	context.JSON(http.StatusCreated,  gin.H{"message": "User registered."})
}

func cancelResgistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, _ := strconv.ParseInt(context.Param("id"), 10, 64)

	event := models.Event{ID: eventId}
	exists, err := event.RegistrationExists(userId)
	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not check registration existence."})
		return
	}

	if !exists {
		context.JSON(http.StatusNotFound, gin.H{"message": "Registration not found for this event and user."})
		return
	}

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration for event"})
		return
	}

	context.JSON(http.StatusOK,  gin.H{"message": "Registration canceled."})
}