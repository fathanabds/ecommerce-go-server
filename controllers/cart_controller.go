package controllers

import (
	"ecommerce-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	// Ambil userID dari context yang di-set oleh middleware
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get userID"})
		return
	}

	// Ambil data dari body request
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validasi apakah produk ada
	var foundProduct *models.Product
	for _, product := range models.Products {
		if product.ID == cartItem.ProductID {
			foundProduct = &product
			break
		}
	}
	if foundProduct == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Tambahkan item ke keranjang user
	uid := userID.(uint)
	models.Carts[uid] = append(models.Carts[uid], cartItem)

	// Kirim respons
	c.JSON(http.StatusOK, gin.H{
		"message": "Product added to cart",
		"cart":    models.Carts[uid],
	})
}

func GetCart(c *gin.Context) {
	// Ambil userID dari context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get userID"})
		return
	}

	// Ambil keranjang berdasarkan userID
	cart, exists := models.Carts[userID.(uint)]
	if !exists || len(cart) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Cart is empty",
			"cart":    []models.CartItem{},
		})
		return
	}

	// Kirim respons
	c.JSON(http.StatusOK, gin.H{
		"message": "Cart fetched successfully",
		"cart":    cart,
	})
}
