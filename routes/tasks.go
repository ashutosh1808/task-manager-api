package routes

import (
	"net/http"
	"strconv"

	"example.com/investment-calculator/practice/task-manager-api/models"
	"github.com/gin-gonic/gin"
)

func getTasks(context *gin.Context) {
	tasks, err := models.GetTasks()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't get tasks", "errors": err.Error()})
		return
	}
	context.JSON(http.StatusOK, tasks)
}

func createTasks(context *gin.Context) {
	var task models.Task
	err := context.ShouldBindJSON(&task)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't bind with json", "errors": err.Error()})
		return
	}

	err = task.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't save data", "errors": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "task created"})
}

func getTask(context *gin.Context) {
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't parse eventId", "errors": err.Error()})
		return
	}

	task, err := models.GetTaskById(taskId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "could not find task", "errors": err.Error()})
		return
	}

	context.JSON(http.StatusOK, task)
}

func updateTask(context *gin.Context) {
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't parse eventId", "errors": err.Error()})
		return
	}

	_, err = models.GetTaskById(taskId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "could not find task", "errors": err.Error()})
		return
	}

	var updatedTask models.Task
	err = context.ShouldBindJSON(&updatedTask)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't parse taskID", "errors": err.Error()})
		return
	}

	updatedTask.ID = taskId
	err = updatedTask.UpdateTask()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update task", "errors": err.Error()})
		return
	}
	context.JSON(http.StatusOK, updatedTask)
}

func deleteTask(context *gin.Context) {
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't parse eventId", "errors": err.Error()})
		return
	}

	event, err := models.GetTaskById(taskId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "could not find task", "errors": err.Error()})
		return
	}

	err = event.DeleteTask()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't parse eventId", "errors": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
