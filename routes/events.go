package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
 
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error() })
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	// context.ShouldBindJSON 将请求体中的 JSON 数据绑定到指定的结构体中
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	
	userId := context.GetInt64("userId")
	event.UserID = userId
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusCreated, gin.H{"message": err.Error() })
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not parse id:%v", err)})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Not found event by id:%v", err.Error() )})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event found successful",
		"event":   event,
	})

}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error() })
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Not found event by id:%v", err) })
		return
	}

	// 验证userId只能由创建者更新
	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusNonAuthoritativeInfo, gin.H{ "message": "Not authorized to update event" })
		return 
	}
	

	var updatedEvent models.Event
	if err = context.ShouldBindJSON(&updatedEvent); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Could not parse request data:%v", err) })
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.UpdateEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error() })
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Update successful"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Not found event by id:%v", err) })
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error() })
		return
	}
	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusNonAuthoritativeInfo, gin.H{ "message": "Not authorized to delete event" })
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error() })
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Delete successful"})
}

