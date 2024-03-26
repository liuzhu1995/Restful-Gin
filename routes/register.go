package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "message": fmt.Sprintf("Could not parse event id:%v", err) })
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": "Could not fetch event." })
		return 
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": "Could not register user." })
		return
	}

	context.JSON(http.StatusOK, gin.H{ "message": "Registered!" })

}
func cancelRegister(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, fmt.Sprintf("Could not parse event id:%v", err))
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": "Could not fetch event." })
		return
	}

	err = event.CancelRegister(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": "Could not cancel register." })
		return
	}
	context.JSON(http.StatusCreated, gin.H{ "message": "Canceled!" })
}


func getRegisterAll(context *gin.Context) {
	registers, err := models.GetAllRegisters()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": err.Error() })
		return 
	}

	context.JSON(http.StatusOK, registers)
}