package config

import (
	"github.com/MRX173/gin-be/internal/models"
)

func SyncDB() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}
