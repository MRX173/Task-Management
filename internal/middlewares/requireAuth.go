package middelwares

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MRX173/gin-be/config"
	"github.com/MRX173/gin-be/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		hmacSecret := []byte("9d76e6d9cb519dbce89d39ff99d4a37aed1a00a83755f9ec7dfda3d596ec7978")

		return hmacSecret, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		var user models.User

		config.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
