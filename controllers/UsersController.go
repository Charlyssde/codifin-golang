package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test.codifin/models"
	"test.codifin/services"
)

type UsersController struct {
	service services.UserService
}

func NewUsersController(service services.UserService) *UsersController {
	return &UsersController{service: service}
}

func (controller *UsersController) Create(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, user)
	}

	res, err := controller.service.Create(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, res)

}

func (controller *UsersController) Login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, user)
	}

	res, err := controller.service.Login(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, res)
}
