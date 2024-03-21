package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error() })
		return
	}

	context.JSON(http.StatusOK, users)
}

func singnup(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error() })
		return
	}

	err := user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error() })
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Created user successful", "user": user})
}

func login(context *gin.Context) {
	var user models.User
	 err := context.ShouldBindJSON(&user)
	 if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })
		return
	}
	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{ "message": err.Error() })
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": fmt.Sprintf("JWT签名失败:%v", err) })
		return
	}

	context.JSON(http.StatusOK, gin.H{ "message": "Login successful", "token": token })
}

func deleteUser(context *gin.Context) {
	userID, err := strconv.ParseInt(context.Param("userID"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "message": err })
		return
	}

	user, err := models.GetUserByID(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": err.Error() })
		return
	}
	err = user.DeleteUser()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{ "message": err.Error() })
		return
	}

	context.JSON(http.StatusOK, gin.H{ "message": "Delete successful" })
}