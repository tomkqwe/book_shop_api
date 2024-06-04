package main

import (
	"book_shop_api/controller"
	"book_shop_api/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Start application...")
	database.DatabaseConnection()

	r := gin.Default()
	r.GET("/books/:id", controller.ReadBook)
	r.GET("/books", controller.ReadBooks)
	r.POST("/books", controller.CreateBook)
	r.PUT("/books/:id", controller.UpdateBook)
	r.DELETE("/books/:id", controller.DeleteBook)

	r.Run(":8080")

}
