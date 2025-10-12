package middleware

import (
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

// DevAuthMiddleware:
// - If DEV_AUTH == "1": reads X-API-User-ID and sets "userID" in context.
// - Otherwise (DEV_AUTH != "1", e.g. "0" or unset): rejects the request with 401.
func DevAuthMiddleware() gin.HandlerFunc {
    enabled := os.Getenv("DEV_AUTH") == "1"

    return func(c *gin.Context) {
        if !enabled {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "DEV_AUTH not enabled"})
            return
        }

        userID := c.GetHeader("X-API-User-ID")
        if userID == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "DEV_AUTH enabled: missing X-API-User-ID header"})
            return
        }

        c.Set("userID", userID)
        c.Next()
    }
}