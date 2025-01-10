package controllers

import (
	"ecommerce-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	// Ambil semua produk
	if len(models.Products) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No products available", "products": []models.Product{}})
		return
	}

	// Kirim respons
	c.JSON(http.StatusOK, gin.H{
		"message":  "Products fetched successfully",
		"products": models.Products,
	})
}

func AddProducts(c *gin.Context) {
	var newProduct models.Product
	err := c.ShouldBindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set ID baru berdasarkan panjang array produk saat ini + 1
	newProduct.ID = uint(len(models.Products) + 1)
	models.Products = append(models.Products, newProduct)
	c.JSON(http.StatusCreated, gin.H{"message": "Product added successfully", "product": newProduct})
}
