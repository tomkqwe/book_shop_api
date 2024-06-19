package http

import (
	"book_shop_api/internal/repository/pgrepo/models"
	"book_shop_api/internal/service"
	"book_shop_api/internal/transport/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartHandlerInterface interface {
	CreateCart(ctx *gin.Context)
	GetCart(ctx *gin.Context)
	GetCarts(ctx *gin.Context)
	UpdateCart(ctx *gin.Context)
	DeleteCart(ctx *gin.Context)
}

type cartHandler struct {
	s service.Service[models.Cart]
}

func NewCartHandler(s service.Service[models.Cart]) CartHandlerInterface {
	return cartHandler{s: s}
}

func (h cartHandler) CreateCart(ctx *gin.Context) {
	var cart models.Cart
	err := ctx.ShouldBind(&cart)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	res, err := h.s.PostEntity(ctx, cart)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"isCrteate": res,
	})
}

func (h cartHandler) GetCart(ctx *gin.Context) {
	ID := ctx.Param("id")
	id, err := uuid.Parse(ID)

	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}
	cart, err := h.s.GetEntity(ctx, id)

	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"cart": cart,
	})
}

func (h cartHandler) GetCarts(ctx *gin.Context) {
	carts, err := h.s.GetEntities(ctx)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"carts": carts,
	})
}

func (h cartHandler) UpdateCart(ctx *gin.Context) {
	var updateCart models.Cart
	idStr := ctx.Param("id")
	err := ctx.ShouldBind(&updateCart)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}
	userID, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	res, err := h.s.PutEntity(ctx, userID, updateCart)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"isUpdated": res,
	})
}

func (h cartHandler) DeleteCart(ctx *gin.Context) {
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
