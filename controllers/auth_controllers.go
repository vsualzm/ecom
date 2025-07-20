package controllers

import (
	"database/sql"
	"ecom/config"
	"ecom/utils"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Input tidak valid", "error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal hash password"})
		return
	}

	query := `INSERT INTO users (username, email, password, role, code_generate) VALUES ($1, $2, $3, $4, $5)`
	_, err = config.DB.Exec(query, input.Username, input.Email, hashedPassword, input.Role, generateRandomCode())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal mendaftarkan user", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Registrasi berhasil"})
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Input tidak valid", "error": err.Error()})
		return
	}

	var userID int
	var hashedPassword string

	err := config.DB.QueryRow("SELECT id, password FROM users WHERE email = $1", input.Email).Scan(&userID, &hashedPassword)
	if err == sql.ErrNoRows || !utils.CheckPasswordHash(input.Password, hashedPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Email atau password salah"})
		return
	}

	token, err := utils.GenerateJWT(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal membuat token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Login berhasil", "data": gin.H{"token": token}})
}

func generateRandomCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 10)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
