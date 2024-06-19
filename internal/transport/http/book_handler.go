package http

import (
	"book_shop_api/internal/repository/pgrepo/models"
	"book_shop_api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookHandlerInterface interface {
	GetBook(c *gin.Context)
	GetBooks(c *gin.Context)
	CreateBook(c *gin.Context)
	UpdateBook(c *gin.Context)
	DeleteBook(c *gin.Context)
}

type bookHandler struct {
	bs service.BookServiceInterface
}

func NewBookHandler(b service.BookServiceInterface) BookHandlerInterface {
	return bookHandler{bs: b}
}

func (bh bookHandler) GetBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	model, err := bh.bs.GetBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": model,
	})
}

func (bh bookHandler) CreateBook(c *gin.Context) {
	var model models.Book
	err := c.ShouldBind(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res, e := bh.bs.CreateBook(model)
	if e != nil {

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": e.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": res,
	})
}

func (bh bookHandler) GetBooks(c *gin.Context) {
	models, err := bh.bs.GetBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"books": models,
	})
}

func (bh bookHandler) UpdateBook(c *gin.Context) {
	var nemModel models.Book
	idStr := c.Param("id")
	e := c.ShouldBind(&nemModel)
	id, err := uuid.Parse(idStr)
	if err != nil || e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error() + e.Error(),
		})
		return
	}
	res, err := bh.bs.UpdateBook(id, nemModel)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"newBook": res,
	})
}

func (bh bookHandler) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := bh.bs.DeleteBook(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isDeleted": res,
	})
}
