package middlewares

import (
	"context"
	"log"

	"github.com/Bobby-P-dev/todo-listgo.git/internal/helpers"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/repository"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

type AuthMiddleware struct {
	UserRepo repository.UserRepo
}

func NewAuthMiddleware(userRepo repository.UserRepo) *AuthMiddleware {
	return &AuthMiddleware{UserRepo: userRepo}
}

func (m *AuthMiddleware) ValidateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		clientId := helpers.GetEnv("CLIENTID")

		token, err := c.Cookie("user-session")
		if err != nil {
			log.Println("[AUTH] Cookie error:", err)
			c.JSON(401, gin.H{"error": "Unauthorized - Cookie missing"})
			c.Abort()
			return
		}

		log.Println("[AUTH] Token received:", token)

		payload, err := idtoken.Validate(context.Background(), token, clientId)
		if err != nil {
			log.Println("[AUTH] Token validation error:", err)
			c.JSON(401, gin.H{"error": "Unauthorized - Invalid token"})
			c.Abort()
			return
		}

		email, ok := payload.Claims["email"].(string)
		if !ok {
			log.Println("[AUTH] Email claim missing")
			c.JSON(401, gin.H{"error": "Unauthorized - Invalid claims"})
			c.Abort()
			return
		}

		log.Println("[AUTH] Email from token:", email)
		data, customErr := m.UserRepo.GetUserByEmail(email)
		if customErr != nil {
			log.Println("[AUTH] DB error:", customErr)
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}

		if data == nil {
			log.Println("[AUTH] User not found for email:", email)
			c.JSON(401, gin.H{"error": "Unauthorized - User not registered"})
			c.Abort()
			return
		}

		log.Println("[AUTH] User authenticated:", data.ID)
		c.Set("user", data)
		c.Next()
	}
}
