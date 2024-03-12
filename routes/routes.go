package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leguzman/go-ecommerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/users/addProduct", controllers.ProductViewAdmin())
	incomingRoutes.POST("/users/productView", controllers.SearchProduct())
	incomingRoutes.POST("/users/search", controllers.SearchProductByQuery())
}
