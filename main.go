package main

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.DB.Close()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.GET("/events/:id", getSingleEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id",func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println(id)
	})
	server.DELETE("/events/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println(id)
	})
	server.POST("/singnup")
	server.POST("/login")
	server.POST("/events/:id/register", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println(id)
	})
	server.DELETE("/events/:id/register", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println(id)
	})
	server.Run()
}


func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	fmt.Println(err, "GetAllEvents")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": fmt.Sprint(err) })
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	// context.ShouldBindJSON 将请求体中的 JSON 数据绑定到指定的结构体中
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	event.UserID = 1
	err := event.Save()
	
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{ "message":  fmt.Sprint(err) })
		return 
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event": event,
	})
}

func getSingleEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "message": fmt.Sprint(err) })
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message":  fmt.Sprint(err)})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event found successed",
		"event": event,
	})

}