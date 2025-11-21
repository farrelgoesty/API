package controller

import (
	"net/http"
	"pos-coffee/config"
	"pos-coffee/model"

	"github.com/gin-gonic/gin"
)

func GetAllMenus(c *gin.Context) {
	query := "SELECT id, kategoriid, namaMenu, harga, stok, image FROM menus"
	rows, err := config.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}
	defer rows.Close()

	var data []model.Menus
	for rows.Next() {
		var m model.Menus
		if err := rows.Scan(&m.ID, &m.Kategori_id, &m.Nama_menu, &m.Harga, &m.Stok, &m.Image); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
			return
		}
		data = append(data, m)
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": data})
}

func GetMenuById(c *gin.Context) {
	id := c.Param("id")
	query := "SELECT id, kategoriid, namaMenu, harga, stok, image FROM menus WHERE id=?"
	row := config.DB.QueryRow(query, id)

	var m model.Menus
	err := row.Scan(&m.ID, &m.Kategori_id, &m.Nama_menu, &m.Harga, &m.Stok, &m.Image)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": 0, "error": "Menu not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": m})
}

func CreateMenu(c *gin.Context) {
	var m model.Menus
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	query := "INSERT INTO menus (kategoriid, namaMenu, harga, stok, image) VALUES (?, ?, ?, ?, ?)"
	_, err := config.DB.Exec(query, m.Kategori_id, m.Nama_menu, m.Harga, m.Stok, m.Image)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Menu added"})
}

func UpdateMenu(c *gin.Context) {
	id := c.Param("id")
	var m model.Menus

	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	query := "UPDATE menus SET kategoriid=?, namaMenu=?, harga=?, stok=?, image=? WHERE id=?"
	_, err := config.DB.Exec(query, m.Kategori_id, m.Nama_menu, m.Harga, m.Stok, m.Image, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Menu updated"})
}

func DeleteMenu(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM menus WHERE id=?"
	_, err := config.DB.Exec(query, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Menu deleted"})
}