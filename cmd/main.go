package main

import (
	"book_shop_api/internal/database"
	"book_shop_api/internal/repository/pgrepo"
	"book_shop_api/internal/service"
	"book_shop_api/internal/transport/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Start application...")
	database.DatabaseConnection()

	r := gin.Default()
	pgRepo := pgrepo.NewBookRepo(database.DB)
	bookSrv := service.NewBookService(pgRepo)
	bookHandler := http.NewBookHandler(bookSrv)
	r.GET("/books/:id", bookHandler.GetBook)
	r.GET("/books", bookHandler.GetBooks)
	r.POST("/books", bookHandler.CreateBook)
	r.PUT("/books/:id", bookHandler.UpdateBook)
	r.DELETE("/books/:id", bookHandler.DeleteBook)

	r.Run(":8080")

}
