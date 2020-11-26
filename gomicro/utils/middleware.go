package utils

import (
	"github.com/gin-gonic/gin"
)


//验证TOKEN中间件
func ValidTokenMiddleware(c *gin.Context) {
	tokenStr := c.GetHeader("Auth")
	_,_,err := ValidateToken(tokenStr)
	if err != nil {
		ResponseError(c,ErrCodeValidToken,nil)
		c.Abort()
		return
	}
	c.Next()
}