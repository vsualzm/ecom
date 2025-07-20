package controllers

import (
	"ecom/config"
	"ecom/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductInput struct {
	Name        string  `json:"name_product"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CodeScan    string  `json:"code_scan"`
	CategoryID  int     `json:"category_id"`
}

type CategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateCategory(c *gin.Context) {
	userID := c.GetInt("user_id")

	var role string
	err := config.DB.QueryRow("SELECT role FROM users WHERE id = $1", userID).Scan(&role)
	if err != nil || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "message": "Hanya admin yang bisa membuat kategori"})
		return
	}

	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Input tidak valid"})
		return
	}

	_, err = config.DB.Exec("INSERT INTO categories (name, description) VALUES ($1, $2)", input.Name, input.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal membuat kategori"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Kategori berhasil dibuat"})
}

func CreateProduct(c *gin.Context) {
	userID := c.GetInt("user_id")

	var role string
	err := config.DB.QueryRow("SELECT role FROM users WHERE id = $1", userID).Scan(&role)
	if err != nil || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "message": "Hanya admin yang bisa membuat produk"})
		return
	}

	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Input tidak valid"})
		return
	}

	_, err = config.DB.Exec(`INSERT INTO products (name_product, description, price, stock, code_scan, category_id) VALUES ($1, $2, $3, $4, $5, $6)`,
		input.Name, input.Description, input.Price, input.Stock, input.CodeScan, input.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal membuat produk"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Produk berhasil dibuat"})
}

func GetProducts(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, name_product, description, price, stock, code_scan, category_id FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal mengambil data produk"})
		return
	}
	defer rows.Close()

	var products []map[string]interface{}
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CodeScan, &p.CategoryID); err != nil {
			continue
		}
		products = append(products, map[string]interface{}{
			"id":           p.ID,
			"name_product": p.Name,
			"description":  p.Description,
			"price":        p.Price,
			"stock":        p.Stock,
			"code_scan":    p.CodeScan,
			"category_id":  p.CategoryID,
		})
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": products})
}

func UpdateProduct(c *gin.Context) {
	userID := c.GetInt("user_id")
	var role string
	err := config.DB.QueryRow("SELECT role FROM users WHERE id = $1", userID).Scan(&role)
	if err != nil || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "message": "Hanya admin yang bisa update produk"})
		return
	}

	id := c.Param("id")
	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Input tidak valid"})
		return
	}

	_, err = config.DB.Exec(`UPDATE products SET name_product=$1, description=$2, price=$3, stock=$4, code_scan=$5, category_id=$6, updated_at=NOW() WHERE id=$7`,
		input.Name, input.Description, input.Price, input.Stock, input.CodeScan, input.CategoryID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal update produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Produk berhasil diupdate"})
}

func DeleteProduct(c *gin.Context) {
	userID := c.GetInt("user_id")
	var role string
	err := config.DB.QueryRow("SELECT role FROM users WHERE id = $1", userID).Scan(&role)
	if err != nil || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "message": "Hanya admin yang bisa menghapus produk"})
		return
	}

	id := c.Param("id")
	_, err = config.DB.Exec("DELETE FROM products WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal hapus produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Produk berhasil dihapus"})
}
