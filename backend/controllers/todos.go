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

// Update a task (Mark as done/undone)
func UpdateTodo(c *gin.Context) {
	// 1. Get the ID from the URL
	id := c.Param("id")

	// 2. Get the data from the body
	var body struct {
		Title     string
		Completed bool
	}
	c.Bind(&body)

	// 3. Find the task in the DB
	var todo models.Todo
	result := initializers.DB.First(&todo, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	// 4. Update it
	initializers.DB.Model(&todo).Updates(models.Todo{
		Title:     body.Title,
		Completed: body.Completed,
	})

	// 5. Respond
	c.JSON(200, gin.H{"todo": todo})
}
