package controller

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/srikanthankamtw/todo-rest/config"
	"github.com/srikanthankamtw/todo-rest/model"
)

func CreateTask(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	var user model.User
	config.GetDB().Where("id = ?", claims[config.IdentityKey]).First(&user)

	if user.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	var todo model.Task
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.UserID = user.ID
	config.GetDB().Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully!", "task": todo})
}

func FetchAllTask(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	var user model.User
	config.GetDB().Where("id = ?", claims[config.IdentityKey]).First(&user)

	if user.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	var todos []model.Task
	config.GetDB().Where("user_id = ?", user.ID).Order("created_at desc").Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No tasks found", "data": todos})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func FetchSingleTask(c *gin.Context) {
	todoID := c.Param("id")

	if len(todoID) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	var todo model.Task
	config.GetDB().First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No todo found!"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func UpdateTask(c *gin.Context) {
	todoID := c.Param("id")

	if len(todoID) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	var newTodo model.Task
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todo model.Task
	config.GetDB().First(&todo, todoID)

	if todo.ID <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No task found"})
		return
	}

	config.GetDB().Model(&todo).Update("title", newTodo.Title)
	config.GetDB().Model(&todo).Update("description", newTodo.Description)

	config.GetDB().First(&todo, todoID)

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully!", "task": todo})
}

func DeleteTask(c *gin.Context) {
	todoID := c.Param("id")

	var todo model.Task

	if len(todoID) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	config.GetDB().First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No task found!"})
		return
	}

	config.GetDB().Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully", "task": todo})
}
