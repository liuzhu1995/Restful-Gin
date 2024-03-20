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
	fmt.Println(err, "GetAllEvents")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprint(err)})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	// context.ShouldBindJSON 将请求体中的 JSON 数据绑定到指定的结构体中
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	event.UserID = 1
	err := event.Save()

	if err != nil {
		context.JSON(http.StatusCreated, gin.H{"message": fmt.Sprint(err)})
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
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("获取数据失败:%v", err)})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event found successed",
		"event":   event,
	})

}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	_, err = models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("此id不存在:%v", err)})
		return
	}

	var event models.Event
	if err = context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("数据格式不匹配:%v", err)})
		return
	}

	event.ID = id
 
	err = event.UpdateEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Update successed"})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("此id不存在:%v", err)})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("无法删除:%v", err)})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Delete successed"})
}

func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": fmt.Sprint(err) })
		return 
	}

	context.JSON(http.StatusOK, users)
}

func createUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "message": fmt.Sprint(err) })
		return
	}

	err := user.Save()

	if err != nil {
		context.JSON(http.StatusCreated, gin.H{ "message": fmt.Sprint(err) })
		return
	}
	context.JSON(http.StatusOK, gin.H{ "message": "Created user successed", "user": user })
}