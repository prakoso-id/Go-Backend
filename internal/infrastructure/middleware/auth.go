package middleware

import (
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-backend/internal/modules/user/domain/repository"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
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

// RateLimiter stores IP-based rate limiters
type RateLimiter struct {
	visitors map[string]*rate.Limiter
	mu      sync.Mutex
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		visitors: make(map[string]*rate.Limiter),
	}
}

// GetLimiter returns a rate limiter for an IP
func (rl *RateLimiter) GetLimiter(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	limiter, exists := rl.visitors[ip]
	if !exists {
		// Allow 100 requests per minute
		limiter = rate.NewLimiter(rate.Every(time.Minute/100), 100)
		rl.visitors[ip] = limiter
	}

	return limiter
}

// RateLimitMiddleware limits the number of requests from an IP
func RateLimitMiddleware(rl *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := rl.GetLimiter(ip)
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, formatResponse(
				http.StatusTooManyRequests,
				"Rate limit exceeded",
				nil,
				"Too many requests. Please try again later.",
			))
			c.Abort()
			return
		}
		c.Next()
	}
}

// TokenType represents the type of token
type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

// JWTAuth middleware for protecting routes
func JWTAuth(tokenType TokenType) gin.HandlerFunc {
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

		tokenString := bearerToken[1]
		claims := jwt.MapClaims{}

		// Validate token signature and claims
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, formatResponse(
				http.StatusUnauthorized,
				"Authentication failed",
				nil,
				"Invalid token signature",
			))
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, formatResponse(
				http.StatusUnauthorized,
				"Authentication failed",
				nil,
				"Invalid token claims",
			))
			c.Abort()
			return
		}

		// Validate token type
		if tokenTypeClaim, ok := claims["token_type"].(string); !ok || TokenType(tokenTypeClaim) != tokenType {
			c.JSON(http.StatusUnauthorized, formatResponse(
				http.StatusUnauthorized,
				"Authentication failed",
				nil,
				"Invalid token type",
			))
			c.Abort()
			return
		}

		// Check token expiration
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				c.JSON(http.StatusUnauthorized, formatResponse(
					http.StatusUnauthorized,
					"Authentication failed",
					nil,
					"Token has expired",
				))
				c.Abort()
				return
			}
		}

		// Get user from database and verify token
		userID := uint(claims["user_id"].(float64))
		db := c.MustGet("db").(*gorm.DB)
		userRepo := repository.NewUserRepository(db)
		user, err := userRepo.GetByID(userID)

		if err != nil {
			c.JSON(http.StatusUnauthorized, formatResponse(
				http.StatusUnauthorized,
				"Authentication failed",
				nil,
				"User not found",
			))
			c.Abort()
			return
		}

		if user.Token == nil || *user.Token != tokenString {
			c.JSON(http.StatusUnauthorized, formatResponse(
				http.StatusUnauthorized,
				"Authentication failed",
				nil,
				"Token has been revoked",
			))
			c.Abort()
			return
		}

		// Store user information in context
		c.Set("user_id", userID)
		c.Set("user_role", claims["role"])
		c.Next()
	}
}

// RequireAuth protects routes that require authentication
func RequireAuth(enabled bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !enabled {
			c.Next()
			return
		}

		_, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, formatResponse(
				http.StatusUnauthorized,
				"Authentication required",
				nil,
				"You must be logged in to access this resource",
			))
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireRole checks if the authenticated user has the required role
func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusForbidden, formatResponse(
				http.StatusForbidden,
				"Access denied",
				nil,
				"Role information not found",
			))
			c.Abort()
			return
		}

		if userRole != role {
			c.JSON(http.StatusForbidden, formatResponse(
				http.StatusForbidden,
				"Access denied",
				nil,
				"Insufficient permissions to access this resource",
			))
			c.Abort()
			return
		}

		c.Next()
	}
}
