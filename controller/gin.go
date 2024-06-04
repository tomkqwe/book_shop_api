package controller

import (
	"book_shop_api/database"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(ctx *gin.Context) {
	var book *database.Book
	err := ctx.ShouldBind(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(book)
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a book",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func ReadBook(ctx *gin.Context) {
	var book database.Book
	id := ctx.Param("id")
	res := database.DB.Find(&book, id)
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"book": book,
		})
	}
}

func ReadBooks(ctx *gin.Context) {
	var books []database.Book
	res := database.DB.Find(&books)
	if res.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("books not found"),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}

func UpdateBook(ctx *gin.Context) {
	var book database.Book
	id := ctx.Param("id")
	err := ctx.ShouldBind(&book)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updateBook database.Book
	res := database.DB.Model(&updateBook).Where("id = ?", id).Updates(book)

	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "book not updated",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"book": book,
		})
	}
}

func DeleteBook(ctx *gin.Context) {
	var book database.Book
	id := ctx.Param("id")
	res := database.DB.Find(&book, id)
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "book not found",
		})
		return
	} else {
		database.DB.Delete(&book)
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("book name: %s , with id %d deleted successfully", book.Title, book.ID),
		})
	}
}
