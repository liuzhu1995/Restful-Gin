package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	fmt.Println(err, "GetAllEvents")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error() })
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	tokenString := context.Request.Header.Get("Authorization")
	fmt.Println("tokenString", tokenString)
	if tokenString == "" {
		context.JSON(http.StatusUnauthorized, gin.H{ "message": "Not authorization" })
		return 
	}

	err := utils.VerifyToken(tokenString)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{ "message": fmt.Sprintf("Not authorization, %v", err) })
		return
	}
	var event models.Event
	// context.ShouldBindJSON 将请求体中的 JSON 数据绑定到指定的结构体中
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	event.UserID = 1
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
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("id解析失败%v", err)})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("获取数据失败:%v", err.Error() )})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event found successful",
		"event":   event,
	})

}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error() })
		return
	}

	_, err = models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("此id不存在:%v", err) })
		return
	}

	var event models.Event
	if err = context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("数据格式不匹配:%v", err) })
		return
	}

	event.ID = id
 
	err = event.UpdateEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error() })
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Update successful"})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error() })
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error() })
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error() })
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Delete successful"})
}

