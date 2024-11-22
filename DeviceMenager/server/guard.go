package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *APIServer) authGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Odczyt ciasteczka
		token, err := c.Cookie(s.config.Server.AuthCookieName)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
			return
		}

		user, ok := s.userHandler.GetUserData(token)
		if ok != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
	
		c.Set("uid", user.ID)
		c.Next()
	}
}