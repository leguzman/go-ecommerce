package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/leguzman/go-ecommerce/controllers"
	"github.com/leguzman/go-ecommerce/database"
	"github.com/leguzman/go-ecommerce/middleware"
	"github.com/leguzman/go-ecommerce/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app := controllers.NewApplication(database.ProdutData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeItem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
