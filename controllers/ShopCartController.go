package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test.codifin/models"
	"test.codifin/services"
)

type ShopCartController struct {
	service services.ShopCartService
}

func NewShopCartController(service services.ShopCartService) *ShopCartController {
	return &ShopCartController{service: service}
}

func (controller *ShopCartController) Create(ctx *gin.Context) {
	var shopCart models.ShopCart
	err := ctx.ShouldBindJSON(&shopCart)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, shopCart)
	}

	res, err := controller.service.Create(shopCart)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, res)
}

func (controller *ShopCartController) GetByUser(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	res, err := controller.service.GetByUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, res)
}
