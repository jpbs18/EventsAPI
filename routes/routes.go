package routes

import (
	"Events-API/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/").Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelResgistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}