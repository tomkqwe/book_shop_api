package http

import (
	"book_shop_api/internal/repository/pgrepo/models"
	"book_shop_api/internal/service"
	"book_shop_api/internal/transport/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderHandlerInterface interface {
	PostOrder(ctx *gin.Context)
	GetOrders(ctx *gin.Context)
	GetOrderById(ctx *gin.Context)
	PutOrderById(ctx *gin.Context)
	DeleteOrderById(ctx *gin.Context)
}

type orderHandler struct {
	s service.Service[models.Order]
}

func NewOrderHandler(s service.Service[models.Order]) OrderHandlerInterface {
	return orderHandler{s: s}
}

func (h orderHandler) PostOrder(ctx *gin.Context) {
	var model models.Order
	err := ctx.ShouldBind(&model)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	cartItem, err := h.s.PostEntity(ctx, model)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"order": cartItem,
	})

}

func (h orderHandler) GetOrders(ctx *gin.Context) {
	cartItems, err := h.s.GetEntities(ctx)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"orders": cartItems,
	})
}

func (h orderHandler) GetOrderById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	cartItems, err := h.s.GetEntity(ctx, id)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"orders": cartItems,
	})
}

func (h orderHandler) PutOrderById(ctx *gin.Context) {
	var updateCartItem models.Order
	idStr := ctx.Param("id")
	err := ctx.ShouldBind(&updateCartItem)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	res, err := h.s.PutEntity(ctx, id, updateCartItem)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"isUpdated": res,
	})
}

func (h orderHandler) DeleteOrderById(ctx *gin.Context) {
	var deleteCartItem models.Order
	idStr := ctx.Param("id")
	err := ctx.ShouldBind(&deleteCartItem)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}
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
