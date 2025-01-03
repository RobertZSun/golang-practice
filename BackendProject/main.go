package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/backend/models"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents) // GET, POST, PUT, DELETE, PATCH
	server.POST("/events", createEvent)

	server.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)

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

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created.",
		"event":   event,
	})

}
