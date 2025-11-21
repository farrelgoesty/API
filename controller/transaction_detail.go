package controller

import (
	"net/http"
	"pos-coffee/config"
	"pos-coffee/model"

	"github.com/gin-gonic/gin"
)

func GetAllTransactionDetails(c *gin.Context) {
	query := "SELECT id, transaksiid, menuId, jumlah, subtotal FROM transactiondetails"
	rows, err := config.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}
	defer rows.Close()

	var data []model.Transactiondetails
	for rows.Next() {
		var td model.Transactiondetails
		if err := rows.Scan(&td.ID, &td.Transaksi_id, &td.Menu_id, &td.Jumlah, &td.Subtotal); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
			return
		}
		data = append(data, td)
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": data})
}

func GetTransactionDetailById(c *gin.Context) {
	id := c.Param("id")
	query := "SELECT id, transaksiid, menuId, jumlah, subtotal FROM transactiondetails WHERE id=?"
	row := config.DB.QueryRow(query, id)

	var td model.Transactiondetails
	err := row.Scan(&td.ID, &td.Transaksi_id, &td.Menu_id, &td.Jumlah, &td.Subtotal)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": 0, "error": "Transaction detail not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": td})
}

func CreateTransactionDetail(c *gin.Context) {
	var td model.Transactiondetails
	if err := c.ShouldBindJSON(&td); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	query := "INSERT INTO transactiondetails (transaksiid, menuId, jumlah, subtotal) VALUES (?, ?, ?, ?)"
	_, err := config.DB.Exec(query, td.Transaksi_id, td.Menu_id, td.Jumlah, td.Subtotal)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Transaction detail added"})
}

func UpdateTransactionDetail(c *gin.Context) {
	id := c.Param("id")
	var td model.Transactiondetails

	if err := c.ShouldBindJSON(&td); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	query := "UPDATE transactiondetails SET transaksiid=?, menuId=?, jumlah=?, subtotal=? WHERE id=?"
	_, err := config.DB.Exec(query, td.Transaksi_id, td.Menu_id, td.Jumlah, td.Subtotal, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Transaction detail updated"})
}

func DeleteTransactionDetail(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM transactiondetails WHERE id=?"
	_, err := config.DB.Exec(query, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Transaction detail deleted"})
}