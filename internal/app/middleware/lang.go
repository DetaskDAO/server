package middleware

import (
	"github.com/gin-gonic/gin"
)

func Lang() gin.HandlerFunc {
	return func(c *gin.Context) {
		// chain
		chain := c.Request.Header.Get("x-lang")
		if chain != "" {
			c.Set("lang", chain)
		} else {
			c.Set("lang", "en") // 默认语言
		}
	}
}
