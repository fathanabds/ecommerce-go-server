package controllers

import (
	"net/http"

	"ecommerce-backend/models"
	"ecommerce-backend/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Fungsi untuk registrasi user
func Register(c *gin.Context) {
	var newUser models.User

	// Validasi input
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newUser.Password = string(hashedPassword)

	// Simpan user ke memory
	newUser.ID = uint(len(models.Users) + 1)
	models.Users = append(models.Users, newUser)

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Fungsi untuk login user
func Login(c *gin.Context) {
	var creds models.Credentials

	// Validasi input
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cari user berdasarkan email
	var foundUser *models.User
	for _, user := range models.Users {
		if user.Email == creds.Email {
			foundUser = &user
			break
		}
	}

	if foundUser == nil || bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(creds.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Buat token JWT
	token, err := utils.GenerateJWT(foundUser.ID, foundUser.Username, foundUser.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
