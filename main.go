package main

import (
	"rest-api-gin-project/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/books", service.GetBooks)
	r.GET("/books/:id", service.GetBook)
	r.GET("/reviews", service.GetReviews)
	r.POST("/books", service.CreateBook)
	r.PATCH("/books/:id", service.UpdateBook)
	r.DELETE("/books/:id", service.DeleteBook)

	r.Run(":8080")
}
