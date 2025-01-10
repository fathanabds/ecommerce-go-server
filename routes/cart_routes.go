package routes

import (
	"ecommerce-backend/controllers"
	"ecommerce-backend/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterCartRoutes: Menambahkan route untuk keranjang belanja
func RegisterCartRoutes(router *gin.Engine) {
	authorized := router.Group("/")
	authorized.Use(middleware.JWTAuthMiddleware())

	authorized.POST("/carts", controllers.AddToCart)
	authorized.GET("/carts", controllers.GetCart)
}
