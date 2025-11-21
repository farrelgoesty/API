package controller

import (
	"net/http"
	"pos-coffee/config"
	"pos-coffee/model"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	query := "SELECT id, namaUser, password, role, username, token FROM users"
	rows, err := config.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.NamaUser, &u.Password, &u.Role, &u.Username, &u.Token); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
			return
		}
		users = append(users, u)
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": users})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	query := "SELECT id, namaUser, password, role, username, token FROM users WHERE id=?"
	row := config.DB.QueryRow(query, id)

	var u model.User
	err := row.Scan(&u.ID, &u.NamaUser, &u.Password, &u.Role, &u.Username, &u.Token)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": 0, "error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": u})
}

func CreateUser(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	query := "INSERT INTO users (namaUser, password, role, username, token) VALUES (?, ?, ?, ?, ?)"
	_, err := config.DB.Exec(query, u.NamaUser, u.Password, u.Role, u.Username, u.Token)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "User added"})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var u model.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	query := "UPDATE users SET namaUser=?, password=?, role=?, username=?, token=? WHERE id=?"
	_, err := config.DB.Exec(query, u.NamaUser, u.Password, u.Role, u.Username, u.Token, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "User updated"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM users WHERE id=?"
	_, err := config.DB.Exec(query, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "User deleted"})
}