package main

import (
	"ecommerce-backend/routes"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	// Memuat file .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	// Register routes
	routes.RegisterUserRoutes(r)
	routes.RegisterProductRoutes(r)
	routes.RegisterCartRoutes(r)

	// Jalankan server
	r.Run(":3000")
}
