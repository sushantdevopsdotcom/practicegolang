package middleware

import (
    "encoding/base64"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if !strings.HasPrefix(authHeader, "Basic ") {
            c.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // Decode the base64-encoded credentials
        encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
        credentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // Split the credentials
        parts := strings.SplitN(string(credentials), ":", 2)
        if len(parts) != 2 {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        username, password := parts[0], parts[1]

        // Replace these with your actual credentials
        if username != "admin" || password != "password" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        c.Next()
    }
}
