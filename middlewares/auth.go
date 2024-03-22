package middlewares

import (
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	tokenString := context.Request.Header.Get("Authorization")

	if tokenString == "" {
		// AbortWithStatusJSON 以指定的http状态码返回JSON格式的响应，并结束后续的处理流程
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ "message": "Not authorization"})
		return
	}

	userId, err := utils.VerifyToken(tokenString)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ "message": "Not authorization"})
		return
	}

	// context.Set 用于在当前的请求上下文中设置一个键值对，将其存储在当前的请求上下文中。使用.GetInt64获取值
	context.Set("userId", userId)
	 
	context.Next()
}
