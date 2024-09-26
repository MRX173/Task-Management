package config

import "github.com/MRX173/gin-be/internal/models"

func syncDB() {
	DB.AutoMigrate(&models.User{})
}
