package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test.codifin/models"
	"test.codifin/services"
)

type ProductsController struct {
	service services.ProductsService
}

func NewProductsController(service services.ProductsService) *ProductsController {
	return &ProductsController{
		service: service,
	}
}

func (controller *ProductsController) Create(ctx *gin.Context) {
	var product models.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, product)
	} else {
		res, err := controller.service.Create(product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			ctx.JSON(http.StatusOK, res)
		}
	}

}

func (controller *ProductsController) Update(ctx *gin.Context) {
	var product models.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, product)
	}
	res, err := controller.service.Update(product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, res)
}

func (controller *ProductsController) Delete(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "El id no corresponde a un número")
	}
	res, err := controller.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, res)
}

func (controller *ProductsController) GetAll(ctx *gin.Context) {
	filter := ctx.Param("filter")
	order := ctx.Param("order")
	if order == "" {
		order = "DESC"
	}
	res, err := controller.service.GetAll(filter, order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, res)
}

func (controller *ProductsController) GetById(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "El id no corresponde a un número")
	}
	res, err := controller.service.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, res)
}
