package routes

import (
	"github.com/gin-gonic/gin"
	"test.codifin/controllers"
)

func ShopCartRoutes(controller *controllers.ShopCartController, router *gin.Engine) *gin.Engine {

	shopCartRouteS := router.Group("/shopcart")

	shopCartRouteS.GET("/:id", controller.GetByUser)
	shopCartRouteS.POST("/add/:scId/:pId", controller.AddProduct)
	shopCartRouteS.POST("/remove/:scId/:pId", controller.RemoveProduct)
	shopCartRouteS.POST("/create", controller.Create)

	return router
}
