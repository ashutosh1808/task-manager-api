package main

import (
	"example.com/investment-calculator/practice/task-manager-api/db"
	"example.com/investment-calculator/practice/task-manager-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	gin.SetMode(gin.DebugMode)

	routes.GetRoutes(server)
	server.Run(":8080")
}
