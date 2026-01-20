package main

import (
	"log"

	"github.com/RoudiOlding/spark-todo/controllers"
	"github.com/RoudiOlding/spark-todo/initializers"
	"github.com/RoudiOlding/spark-todo/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// This runs BEFORE main()
	// 1. Load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. Connect to Database
	initializers.ConnectToDB()

	// 3. Auto-Migrate (The Magic Step) 🪄
	initializers.DB.AutoMigrate(&models.Todo{})
}

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  "Sparks Studio Backend is Online",
		})
	})

	r.POST("/todos", controllers.CreateTodo) // Create
	r.GET("/todos", controllers.GetTodos)    // Read

	r.Run(":8080")
}
