// Generic, like a form.

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

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, assuming cloud enviroment")
	}

	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.Todo{})
}

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://spark-todo.vercel.app"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  "Sparks Studio Backend is Online",
		})
	})

	r.POST("/todos", controllers.CreateTodo)
	r.GET("/todos", controllers.GetTodos)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)

	r.Run(":8080")
}
