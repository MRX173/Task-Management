package main

import (
	"net/http"

	"github.com/MRX173/gin-be/config"
	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectDB()
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin!",
		})
	})

	router.Run(":8080")
}
