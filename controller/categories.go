package controller

import (
	"net/http"
	"pos-coffee/config"
	"pos-coffee/model"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	query := "SELECT id, namaKategori FROM categories"
	rows, err := config.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}
	defer rows.Close()

	var data []model.Categories
	for rows.Next() {
		var ct model.Categories
		if err := rows.Scan(&ct.ID, &ct.Nama_kategori); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
			return
		}
		data = append(data, ct)
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": data})
}

func GetCategoryById(c *gin.Context) {
	id := c.Param("id")
	query := "SELECT id, namaKategori FROM categories WHERE id=?"
	row := config.DB.QueryRow(query, id)

	var ct model.Categories
	err := row.Scan(&ct.ID, &ct.Nama_kategori)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": 0, "error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": ct})
}

func CreateCategory(c *gin.Context) {
	var ct model.Categories
	if err := c.ShouldBindJSON(&ct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	query := "INSERT INTO categories (namaKategori) VALUES (?)"
	_, err := config.DB.Exec(query, ct.Nama_kategori)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Category added"})
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var ct model.Categories

	if err := c.ShouldBindJSON(&ct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	query := "UPDATE categories SET namaKategori=? WHERE id=?"
	_, err := config.DB.Exec(query, ct.Nama_kategori, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Category updated"})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM categories WHERE id=?"
	_, err := config.DB.Exec(query, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Category deleted"})
}