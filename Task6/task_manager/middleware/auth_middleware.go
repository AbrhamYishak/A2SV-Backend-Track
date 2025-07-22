package middleware

import (
	"net/http"
	"strings"
    "task6/data"
	"github.com/gin-gonic/gin"
)
func JwtAuthMiddlewareUser(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, _ := data.VerifyToken(authToken, secret)
			if authorized {
				_,err := data.ExtractFromToken(authToken, secret)
				if err != nil {
					c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":"could not extract username from token"})
					c.Abort()
					return
				}
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"message":"invalid token"})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"message":"invalid authHeader"})
		c.Abort()
	}
}
func JwtAuthMiddlewareAdmin(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, _ := data.VerifyToken(authToken, secret)
			if authorized {
				username, err := data.ExtractFromToken(authToken, secret)
				if err != nil {
					c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":"could not extract username from token"})
					c.Abort()
					return
				}
				isadmin,err := data.Isadmin(username)
				if err != nil {
					c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":err.Error()})
					c.Abort()
					return
				}
				if !isadmin{
					c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":"access only for admin"})
					c.Abort()
					return
				}
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"message":"invalid token"})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"message":"invalid authHeader"})
		c.Abort()
	}
}
