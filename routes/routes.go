package routes

import (
	"github.com/gin-gonic/gin"
	"test.codifin/controllers"
)

func ProductRoutes(controller *controllers.ProductsController, router *gin.Engine) *gin.Engine {
	productsRoutes := router.Group("/products")
	productsRoutes.GET("/:id", controller.GetById)
	productsRoutes.GET("/all/:filter/:order", controller.GetAll)
	productsRoutes.POST("", controller.Create)
	productsRoutes.PATCH("", controller.Update)
	productsRoutes.DELETE("/:id", controller.Delete)

	return router
}
