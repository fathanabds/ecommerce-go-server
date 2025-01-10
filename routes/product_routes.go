package routes

import (
	"ecommerce-backend/controllers"
	"ecommerce-backend/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterProductRoutes: Menambahkan route untuk produk
func RegisterProductRoutes(router *gin.Engine) {
	authorized := router.Group("/")
	authorized.Use(middleware.JWTAuthMiddleware())

	authorized.GET("/products", controllers.GetProducts)
	authorized.POST("/products", middleware.RoleAuthMiddleware("admin"), controllers.AddProducts)
}
