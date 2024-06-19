package http

import (
	"book_shop_api/internal/repository/pgrepo/models"
	"book_shop_api/internal/service"
	"book_shop_api/internal/transport/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartItemHandlerInterface interface {
	PostCartItem(ctx *gin.Context)
	GetCartItems(ctx *gin.Context)
	GetCartItemById(ctx *gin.Context)
	PutCartItemById(ctx *gin.Context)
	DeleteCartItemById(ctx *gin.Context)
}

type cartItemHandler struct {
	s service.Service[models.CartItem]
}

func NewCartItemHandler(s service.Service[models.CartItem]) CartItemHandlerInterface {
	return cartItemHandler{s: s}
}

func (h cartItemHandler) PostCartItem(ctx *gin.Context) {
	var model models.CartItem
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
		"cartItem": cartItem,
	})
}

func (h cartItemHandler) GetCartItems(ctx *gin.Context) {
	cartItems, err := h.s.GetEntities(ctx)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"cartItems": cartItems,
	})
}

func (h cartItemHandler) GetCartItemById(ctx *gin.Context) {
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
		"cartItems": cartItems,
	})
}

func (h cartItemHandler) PutCartItemById(ctx *gin.Context) {
	var updateCartItem models.CartItem
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

func (h cartItemHandler) DeleteCartItemById(ctx *gin.Context) {
	var deleteCartItem models.CartItem
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
