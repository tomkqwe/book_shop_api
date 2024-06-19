package http

import (
	"book_shop_api/internal/repository/pgrepo/models"
	"book_shop_api/internal/service"
	"book_shop_api/internal/transport/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookHandlerInterface interface {
	GetBook(ctx *gin.Context)
	GetBooks(ctx *gin.Context)
	CreateBook(ctx *gin.Context)
	UpdateBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
}

type bookHandler struct {
	s service.Service[models.Book]
}

func NewBookHandler(s service.Service[models.Book]) BookHandlerInterface {
	return bookHandler{s: s}
}

func (h bookHandler) GetBook(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	model, err := h.s.GetEntity(ctx, id)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": model,
	})
}

func (h bookHandler) CreateBook(ctx *gin.Context) {
	var model models.Book
	err := ctx.ShouldBind(&model)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}
	res, err := h.s.PostEntity(ctx, model)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusUnprocessableEntity)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": res,
	})
}

func (h bookHandler) GetBooks(ctx *gin.Context) {
	models, err := h.s.GetEntities(ctx)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"books": models,
	})
}

func (h bookHandler) UpdateBook(ctx *gin.Context) {
	var nemModel models.Book
	idStr := ctx.Param("id")
	err := ctx.ShouldBind(&nemModel)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	res, err := h.s.PutEntity(ctx, id, nemModel)

	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"newBook": res,
	})
}

func (h bookHandler) DeleteBook(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	res, err := h.s.DeleteEntity(ctx, id)

	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"isDeleted": res,
	})
}
