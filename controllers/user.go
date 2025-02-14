package controllers

import (
	"book-management-api/database"
	"book-management-api/repository"
	"book-management-api/structs"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser - Mendaftarkan user baru
func RegisterUser(c *gin.Context) {
	var user structs.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah username sudah digunakan
	existingUser, err := repository.GetUserByUsername(database.DbConnection, user.Username)
	if err == nil && existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Set nilai default untuk CreatedAt & CreatedBy
	user.CreatedBy = user.Username
	user.CreatedAt = time.Now()

	// Simpan user ke database
	if err := repository.CreateUser(database.DbConnection, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login - Autentikasi user & generate JWT token
func Login(c *gin.Context) {
	var loginReq structs.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil user dari database berdasarkan username
	user, err := repository.GetUserByUsername(database.DbConnection, loginReq.Username)
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials (user not found)"})
		return
	}

	// Debug: Print password dari database dan input user
	fmt.Println("Stored Hash:", user.Password)
	fmt.Println("Entered Password:", loginReq.Password)

	// Bandingkan password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials (wrong password)"})
		return
	}

	// Generate JWT token
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
