package controllers

import (
	"encoding/json"
	"fmt"
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
	if err := json.NewDecoder(ctx.Request.Body).Decode(&shopCart); err != nil {
		ctx.JSON(http.StatusBadRequest, shopCart)
	} else {
		fmt.Println(shopCart)
		res, err := controller.service.Create(shopCart)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			ctx.JSON(http.StatusOK, res)
		}
	}
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

func (controller *ShopCartController) AddProduct(ctx *gin.Context) {
	scId := ctx.Param("scId")
	pId := ctx.Param("pId")

	id, err := strconv.Atoi(scId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	productId, err := strconv.Atoi(pId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	res, err := controller.service.AddProduct(id, productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, res)
}

func (controller *ShopCartController) RemoveProduct(ctx *gin.Context) {
	scId := ctx.Param("scId")
	pId := ctx.Param("pId")

	id, err := strconv.Atoi(scId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	productId, err := strconv.Atoi(pId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	res, err := controller.service.RemoveProduct(id, productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, res)
}
