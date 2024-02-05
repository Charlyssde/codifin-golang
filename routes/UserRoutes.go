package routes

import (
	"github.com/gin-gonic/gin"
	"test.codifin/controllers"
)

func UserRoutes(controller *controllers.UsersController, router *gin.Engine) *gin.Engine {
	userRoutes := router.Group("/user")
	userRoutes.POST("", controller.Create)
	userRoutes.POST("/login", controller.Login)

	return router
}
