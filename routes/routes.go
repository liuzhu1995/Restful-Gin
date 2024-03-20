package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/singnup", createUser)
	server.POST("/login")
	server.POST("/events/:id/register", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println(id)
	})
	server.DELETE("/events/:id/register", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println(id)
	})

	server.GET("/users", getUsers)
}