package middleware

import (
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	// write custom logic
	return func(c *gin.Context) {
		if !(c.Request.Header.Get("Token") == "auth") {
			c.AbortWithStatusJSON(500, gin.H{
				"Message": "token is required",
			})
			return
		}
		c.Next()
	}
}

// func Authenticate(c *gin.Context) {
// 	if !(c.Request.Header.Get("Token")=="auth"){
// 		c.AbortWithStatusJSON(500,gin.H{
// 			"Message": "token is required",
// 		})
// 		return
// 	}
// 	c.Next()
// }

func AddHeader(c *gin.Context) {
	c.Writer.Header().Set("Key", "value")
	c.Next()
}
