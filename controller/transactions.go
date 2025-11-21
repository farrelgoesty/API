package controller

import (
	"net/http"
	"pos-coffee/config"
	"pos-coffee/model"

	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context) {
	query := "SELECT id, userId, namaCust, tanggal, totalHarga, metodeBayar FROM transactions"
	rows, err := config.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}
	defer rows.Close()

	var data []model.Transactions
	for rows.Next() {
		var t model.Transactions
		if err := rows.Scan(&t.ID, &t.User_id, &t.Nama_cust, &t.Tanggal, &t.Total_harga, &t.Metode_bayar); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
			return
		}
		data = append(data, t)
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": data})
}

func GetTransactionById(c *gin.Context) {
	id := c.Param("id")
	query := "SELECT id, userId, namaCust, tanggal, totalHarga, metodeBayar FROM transactions WHERE id=?"
	row := config.DB.QueryRow(query, id)

	var t model.Transactions
	err := row.Scan(&t.ID, &t.User_id, &t.Nama_cust, &t.Tanggal, &t.Total_harga, &t.Metode_bayar)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": 0, "error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": t})
}

func CreateTransaction(c *gin.Context) {
	var t model.Transactions
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	query := "INSERT INTO transactions (userId, namaCust, tanggal, totalHarga, metodeBayar) VALUES (?, ?, ?, ?, ?)"
	_, err := config.DB.Exec(query, t.User_id, t.Nama_cust, t.Tanggal, t.Total_harga, t.Metode_bayar)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Transaction added"})
}

func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	var t model.Transactions

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	query := "UPDATE transactions SET userId=?, namaCust=?, tanggal=?, totalHarga=?, metodeBayar=? WHERE id=?"
	_, err := config.DB.Exec(query, t.User_id, t.Nama_cust, t.Tanggal, t.Total_harga, t.Metode_bayar, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Transaction updated"})
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM transactions WHERE id=?"
	_, err := config.DB.Exec(query, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Transaction deleted"})
}