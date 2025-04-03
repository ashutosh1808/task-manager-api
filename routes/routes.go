package routes

import "github.com/gin-gonic/gin"

func GetRoutes(server *gin.Engine) {
	server.GET("/tasks", getTasks)
	server.POST("/tasks", createTasks)
	server.GET("/tasks/:id", getTask)
	server.PUT("/tasks/:id", updateTask)
	server.DELETE("/tasks/:id", deleteTask)
}
