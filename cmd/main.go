package main

import (
	"net/http"

	"github.com/MRX173/gin-be/config"
	controller "github.com/MRX173/gin-be/internal/controllers"
	middelwares "github.com/MRX173/gin-be/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectDB()
	config.SyncDB()
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin!",
		})
	})

	router.POST("/signup", controller.Signup)
	router.POST("/login", controller.Login)
	router.GET("/validate", middelwares.RequireAuth, controller.ValidateToken)

	router.Run(":8080")
}
