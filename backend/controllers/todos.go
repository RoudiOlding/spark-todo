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
		Title string `json:"title" binding:"required,min=3"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{"error": "Task title must be at least 3 characters long"})
		return
	}

	// 2. Create the todo
	todo := models.Todo{Title: body.Title, Completed: false}
	result := initializers.DB.Create(&todo)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
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
	initializers.DB.Order("created_at desc").Find(&todos)

	// 2. Return them
	c.JSON(200, gin.H{
		"todos": todos,
	})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title     string
		Completed bool
	}
	c.Bind(&body)

	var todo models.Todo
	result := initializers.DB.First(&todo, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	updateData := map[string]interface{}{
		"completed": body.Completed,
	}

	if body.Title != "" {
		updateData["title"] = body.Title
	}

	initializers.DB.Model(&todo).Updates(updateData)

	c.JSON(200, gin.H{"todo": todo})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Todo{}, id)

	c.Status(200)
}
