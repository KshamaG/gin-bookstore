package main

import (
	"net/http"

	"github.com/KshamaG/gin-bookstore/controllers"
	"github.com/KshamaG/gin-bookstore/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world from bookstore"})
	})

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// Run the server
	r.Run(":8085")
}
