package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"test.codifin/config"
	"test.codifin/controllers"
	"test.codifin/data"
	routes2 "test.codifin/routes"
	"test.codifin/services"
)

func main() {

	db := config.DBConnection()

	routes := gin.Default()

	startServices(db, routes)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func startServices(db *gorm.DB, routes *gin.Engine) {
	productsRepository := data.NewProductImpl(db)
	productsService := services.NewProductServiceImpl(productsRepository)
	productsController := controllers.NewProductsController(productsService)

	usersRepository := data.NewUserRepositoryImpl(db)
	usersService := services.NewUserServiceImpl(usersRepository)
	usersControllers := controllers.NewUsersController(usersService)

	shopCartRepository := data.NewShopCartRepositoryImpl(db)
	shopCartService := services.NewShopCartServiceImpl(shopCartRepository)
	shopCartControllers := controllers.NewShopCartController(shopCartService)

	routes2.ShopCartRoutes(shopCartControllers, routes)
	routes2.UserRoutes(usersControllers, routes)
	routes2.ProductRoutes(productsController, routes)
}
