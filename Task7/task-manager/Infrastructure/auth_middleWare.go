package Infrastructure

import (
	"net/http"
	"strings"
    "task7/Usecases"
	"github.com/gin-gonic/gin"
)
func JwtAuthMiddlewareUser(uc Usecases.UserUsecase ) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, _ := uc.TokenService.VerifyToken(authToken)
			if authorized {
				_,err := uc.TokenService.ExtractFromToken(authToken)
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
func JwtAuthMiddlewareAdmin(uc Usecases.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, _ := uc.TokenService.VerifyToken(authToken) 
			if authorized {
				username, err := uc.TokenService.ExtractFromToken(authToken)
				if err != nil {
					c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":"could not extract username from token"})
					c.Abort()
					return
				}
				isadmin,err := uc.Repo.Isadmin(username)
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
