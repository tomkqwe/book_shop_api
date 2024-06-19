package http

import (
	"book_shop_api/internal/repository/pgrepo/models"
	"book_shop_api/internal/service"
	"book_shop_api/internal/transport/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderItemInterface interface {
	GetOrderItem(ctx *gin.Context)
	GetOrderItems(ctx *gin.Context)
	CreateOrderItem(ctx *gin.Context)
	UpdateOrderItem(ctx *gin.Context)
	DeleteOrderItem(ctx *gin.Context)
}

type orderItemHandler struct {
	s service.Service[models.OrderItem]
}

func NewOrderItemHandler(s service.Service[models.OrderItem]) OrderItemInterface {
	return orderItemHandler{s: s}
}

func (h orderItemHandler) GetOrderItem(ctx *gin.Context) {
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
		"order_item": model,
	})
}

func (h orderItemHandler) GetOrderItems(ctx *gin.Context) {
	models, err := h.s.GetEntities(ctx)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"order_items": models,
	})
}

func (h orderItemHandler) CreateOrderItem(ctx *gin.Context) {
	var model models.OrderItem
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
		"order_item": res,
	})
}

func (h orderItemHandler) UpdateOrderItem(ctx *gin.Context) {
	var nemModel models.OrderItem
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
		"success": res,
	})
}

func (h orderItemHandler) DeleteOrderItem(ctx *gin.Context) {
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
