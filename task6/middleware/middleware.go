package middleware

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_secret_key")

func Authenticate() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")

        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
            c.Abort()
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
            c.Abort()
            return
        }

        tokenString := parts[1]

        claims := jwt.MapClaims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
            c.Abort()
            return
        }

        userID, ok := claims["user_id"].(string)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in token"})
            c.Abort()
            return
        }

        username, ok := claims["username"].(string)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "username not found in token"})
            c.Abort()
            return
        }

        role, ok := claims["role"].(string)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "role not found in token"})
            c.Abort()
            return
        }

        c.Set("userId", userID)
        c.Set("username", username)
        c.Set("role", role)

        c.Next()
    }
}
func Authorize(requiredRole string) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Retrieve the role from the context set by the Authenticate middleware
        role, exists := c.Get("role")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            c.Abort()
            return
        }

        // Check if the user's role matches the required role
        if role != requiredRole {
            c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: insufficient permissions"})
            c.Abort()
            return
        }

        c.Next()
    }
}
