package controllers

import (
	"github.com/RoudiOlding/spark-todo/initializers"
	"github.com/RoudiOlding/spark-todo/models"
	"github.com/gin-gonic/gin"
)

// POST /todos
// Create a new task
func CreateTodo(c *gin.Context) {
	// 1. Get data off request body
	var body struct {
		Title string
	}

	c.Bind(&body)

	// 2. Create the todo
	todo := models.Todo{Title: body.Title, Completed: false}
	result := initializers.DB.Create(&todo) // Pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	// 3. Return it
	c.JSON(200, gin.H{
		"todo": todo,
	})
}

// GET /todos
// Get all tasks
func GetTodos(c *gin.Context) {
	// 1. Get the todos
	var todos []models.Todo
	initializers.DB.Find(&todos)

	// 2. Return them
	c.JSON(200, gin.H{
		"todos": todos,
	})
}
