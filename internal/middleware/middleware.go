package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthRequired scenes software that allows request to communicate and interact with application by authentication.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var signatureKey = os.Getenv("JWT_SIGNATURE_KEY")
		authorizationHeader := c.GetHeader("Authorization")

		if authorizationHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "forbidden",
				"error":   "Token not found",
			})
			return
		} else if !strings.Contains(authorizationHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "forbidden",
				"error":   "invalid token, need bearer token",
			})
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(signatureKey), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
				"error":   err.Error(),
			})
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		t := time.Now()

		userName := claims["user_name"]
		name := claims["name"]
		role := claims["role"]
		timestamp := claims["timestamp"]

		c.Set("userName", userName)
		c.Set("name", name)
		c.Set("role", role)
		c.Set("timestamp", timestamp)
		c.Set("claims", claims)
		// before request
		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
