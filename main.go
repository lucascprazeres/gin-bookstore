package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucascprazeres/gin-bookstore/controllers"
	"github.com/lucascprazeres/gin-bookstore/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	err := r.Run()

	if err != nil {
		return
	}
}
