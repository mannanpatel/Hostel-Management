package middleware

import (
	"fmt"
	"hst_manag/internal/utils"
	gen "hst_manag/internal/utils/generic"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			fmt.Println("========Authorization header missing=============11=====>>")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gen.HandleError(nil, "Bearer token is required"))
			c.Abort()
			return
		}
		fmt.Println("==========Token String from Header:========22========>>>", tokenString)

		claims, err := utils.ValidateToken(tokenString)
		fmt.Println("=========Token validation error:=====33===========>>", claims)
		if err != nil {
			fmt.Println("=========Token validation error:=====35===========>>", err)
			c.JSON(http.StatusUnauthorized, gen.HandleError(err, "Invalid token"))
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("password", claims.FirstName)

		c.Next()
	}
}
