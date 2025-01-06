package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func formatResponse(status int, message string, data interface{}, err string) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, formatResponse(
				http.StatusUnauthorized,
				"Authentication failed",
				nil,
				"Authorization header is required",
			))
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, formatResponse(
				http.StatusUnauthorized,
				"Authentication failed",
				nil,
				"Invalid token format. Use 'Bearer <token>'",
			))
			c.Abort()
			return
		}

		token := bearerToken[1]
		claims := jwt.MapClaims{}

		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !parsedToken.Valid {
			c.JSON(http.StatusUnauthorized, formatResponse(
				http.StatusUnauthorized,
				"Authentication failed",
				nil,
				"Invalid or expired token",
			))
			c.Abort()
			return
		}

		userID := uint(claims["user_id"].(float64))
		c.Set("user_id", userID)
		c.Next()
	}
}
