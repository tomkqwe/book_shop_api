package http

import (
	"book_shop_api/internal/repository/pgrepo/models"
	"book_shop_api/internal/service"
	"book_shop_api/internal/transport/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandlerInterface interface {
	//admin Only
	GetUsers(ctx *gin.Context)
	//admin Only
	GetUserById(ctx *gin.Context)
	//admin Only
	PostUser(ctx *gin.Context)
	PutUser(ctx *gin.Context)
	DeleteUserById(ctx *gin.Context)
}

type userHandler struct {
	s service.Service[models.User]
}

func NewUserHandler(s service.Service[models.User]) UserHandlerInterface {
	return userHandler{s: s}
}

func (h userHandler) GetUsers(ctx *gin.Context) {
	users, err := h.s.GetEntities(ctx)

	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (h userHandler) GetUserById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	user, err := h.s.GetEntity(ctx, id)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (h userHandler) PostUser(ctx *gin.Context) {
	var model models.User
	err := ctx.ShouldBind(&model)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}
	res, err := h.s.PostEntity(ctx, model)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": res,
	})
}

func (h userHandler) PutUser(ctx *gin.Context) {
	var updateUser models.User
	idStr := ctx.Param("id")
	err := ctx.ShouldBind(&updateUser)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	res, err := h.s.PutEntity(ctx, id, updateUser)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"isUpdated": res,
	})
}

func (h userHandler) DeleteUserById(ctx *gin.Context) {
	var deleteUser models.User
	idStr := ctx.Param("id")
	err := ctx.ShouldBind(&deleteUser)
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
