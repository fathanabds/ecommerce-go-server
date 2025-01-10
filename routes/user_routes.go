package routes

import (
	"ecommerce-backend/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes: Menambahkan route untuk user
func RegisterUserRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
}
