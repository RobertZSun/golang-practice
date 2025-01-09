package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/backend/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) // GET, POST, PUT, DELETE, PATCH
	server.GET("/events/:id", getEventByID)

	authenticatedGroup := server.Group("/")
	authenticatedGroup.Use(middlewares.Authenticate)
	authenticatedGroup.POST("/events", createEvent)
	authenticatedGroup.PUT("/events/:id", updateEventByID)
	authenticatedGroup.DELETE("/events/:id", deleteEvent)
	authenticatedGroup.POST("/events/:id/register", registerForEvent)
	authenticatedGroup.DELETE("/events/:id/register", cancelRegistration)

	// server.POST("/events", createEvent)
	// server.PUT("/events/:id", updateEventByID)
	// server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signupUser)
	server.POST("/login", loginUser)
}
