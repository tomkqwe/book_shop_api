package http

import (
	"book_shop_api/internal/repository/pgrepo/models"
	"book_shop_api/internal/service"
	"book_shop_api/internal/transport/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type authHandler struct {
	s service.UserService
}

func NewRegisterHandler(s service.UserService) authHandler {
	return authHandler{s: s}
}

func (h authHandler) Register(ctx *gin.Context) {
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
	claims := &utils.Claims{
		Id:    res.ID,
		Email: res.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("6b60e7ada6a1571fc2473150c9283235d69062f7c32891434af424075d9277a9"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":  res,
		"token": tokenString,
	})
}

func (h authHandler) Login(ctx *gin.Context) {
	var model models.User
	err := ctx.ShouldBindJSON(&model)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	res, err := h.s.GetUserByEmail(ctx, model.Email)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	if !utils.CheckPasswordHash(model.Password, res.Password) {
		utils.WriteErrorResponse(ctx, fmt.Errorf("invalid email or password"), http.StatusUnauthorized)
		return
	}

	claims := &utils.Claims{
		Id: res.ID, // Используем ID пользователя для токена
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("6b60e7ada6a1571fc2473150c9283235d69062f7c32891434af424075d9277a9"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}

	// Возвращаем пользователя и токен в ответе
	ctx.JSON(http.StatusOK, gin.H{
		"user":  res,
		"token": tokenString,
	})
}
