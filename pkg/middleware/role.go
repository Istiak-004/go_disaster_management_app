// middleware/role.go
package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Middleware to allow only admins to access certain routes
func AdminOnly() gin.HandlerFunc {
    return func(c *gin.Context) {
        role := c.GetString("role")
        if role != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Admin access only"})
            c.Abort()
            return
        }
        c.Next()
    }
}

// Middleware to allow only volunteers to access certain routes
func VolunteerOnly() gin.HandlerFunc {
    return func(c *gin.Context) {
        role := c.GetString("role")
        isVerified := c.GetString("isVerified")
        if role != "volunteer" && isVerified !="unverified" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Volunteer access only"})
            c.Abort()
            return
        }
        c.Next()
    }
}
