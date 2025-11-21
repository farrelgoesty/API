package main

import (
	"pos-coffee/config"
	"pos-coffee/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Konek()

	r := gin.Default()

	r.GET("/users", controller.GetAllUsers)
	r.GET("/transaksi", controller.GetAllTransactions)
	r.GET("/detail",controller.GetAllTransactionDetails)
	r.GET("/menu", controller.GetAllMenus)
	r.GET("/kategori", controller.GetAllCategories)

	r.Run(":8000") 
}