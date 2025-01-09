package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/backend/models"
)

func getEvents(context *gin.Context) {
	events, error := models.GetAllEvents()

	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get events.",
			"error":   error,
		})
		return
	}

	context.JSON(http.StatusOK, events)

}

func getEventByID(context *gin.Context) {
	eventId, parseError := strconv.ParseInt(context.Param("id"), 10, 64)

	if parseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse id.",
			"error":   parseError,
		})
		return
	}

	event, error := models.GetEventByID(eventId)

	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event.",
			"error":   error,
		})
		return
	}

	context.JSON(http.StatusOK, event)

}

func createEvent(context *gin.Context) {

	var event models.Event

	// body, _ := io.ReadAll(context.Request.Body)
	// fmt.Println(string(body))

	error := context.ShouldBindJSON(&event)

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
			"error":   error,
		})
		return
	}

	userId := context.GetInt64("userId")
	// event.ID = 1
	event.UserID = userId

	result, SaveErr := event.Save()

	if SaveErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save event.",
			"error":   SaveErr,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created.",
		"event":   result,
	})

}

func updateEventByID(context *gin.Context) {
	eventId, parseError := strconv.ParseInt(context.Param("id"), 10, 64)

	if parseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
			"error":   parseError,
		})
		return
	}

	userId := context.GetInt64("userId")
	event, fetchError := models.GetEventByID(eventId)

	if fetchError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event.",
			"error":   fetchError,
		})
		return
	}

	if userId != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized to update event.",
		})
		return
	}

	var newEvent models.Event

	bindDataErr := context.ShouldBindJSON(&newEvent)

	if bindDataErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not parse event data.",
			"error":   bindDataErr,
		})
		return
	}

	newEvent.ID = eventId

	res, updateErr := newEvent.UpdateEventByID()

	if updateErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event.",
			"error":   updateErr,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated.",
		"result":  res,
	})

}

func deleteEvent(context *gin.Context) {
	eventId, parseError := strconv.ParseInt(context.Param("id"), 10, 64)

	if parseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
			"error":   parseError,
		})
		return
	}

	userId := context.GetInt64("userId")
	event, fetchError := models.GetEventByID(eventId)

	if fetchError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "The event does not exist at all.",
			"error":   fetchError,
		})
		return
	}

	if userId != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized to delete event.",
		})
		return
	}

	res, deleteErr := event.DeleteEventByID()

	if deleteErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete event.",
			"error":   deleteErr,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted.",
		"result":  res,
	})
}
