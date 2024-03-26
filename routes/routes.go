package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// 使用中间件 
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegister)
	authenticated.GET("/events/registers", getRegisterAll)


	server.POST("/singnup", singnup)
	server.POST("/login", login)
	 
	server.GET("/users", getUsers)
	server.DELETE("/users/:userID", deleteUser)
}


// func RegisterRoutes(server *gin.Engine) {
// 	server.GET("/events", getEvents)
// 	server.GET("/events/:id", getEvent)
// // 两种使用中间件的方式，其一
// 	server.POST("/events", middlewares.Authenticate, createEvent)
// 	server.PUT("/events/:id", middlewares.Authenticate, updateEvent)
// 	server.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)


// 	server.POST("/singnup", singnup)
// 	server.POST("/login", login)
// 	server.POST("/events/:id/register", func(c *gin.Context) {
// 		id := c.Param("id")
// 		fmt.Println(id)
// 	})
// 	server.DELETE("/events/:id/register", func(c *gin.Context) {
// 		id := c.Param("id")
// 		fmt.Println(id)
// 	})

// 	server.GET("/users", getUsers)
// 	server.DELETE("/users/:userID", deleteUser)
// }