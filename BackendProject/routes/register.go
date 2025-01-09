package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/backend/models"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, parseError := strconv.ParseInt(context.Param("id"), 10, 64)
	if parseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse id.",
			"error":   parseError,
		})
		return
	}

	event, fetchError := models.GetEventByID(eventId)

	if fetchError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event.",
			"error":   fetchError,
		})
		return
	}

	registerError := event.Register(userId)

	if registerError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not register for event.",
			"error":   registerError,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Successfully registered for event.",
	})

}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, parseError := strconv.ParseInt(context.Param("id"), 10, 64)

	if parseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse id.",
			"error":   parseError,
		})
		return
	}

	var event models.Event
	event.ID = eventId
	cancelError := event.Cancel(userId)

	if cancelError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not cancel for the event.",
			"error":   cancelError,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully canceled the event.",
	})
}
