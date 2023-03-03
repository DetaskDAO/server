package middleware

import (
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/model/response"
	"github.com/gin-gonic/gin"
)

func Chain() gin.HandlerFunc {
	return func(c *gin.Context) {
		// chain
		chain := c.Request.Header.Get("x-chain")
		if chain != "" {
			// 判断链合法性
			_, ok := global.ProviderMap[chain]
			if !ok {
				response.FailWithDetailed(gin.H{"reload": true}, "非法访问", c)
				c.Abort()
				return
			}
			c.Set("chain", chain)
		} else {
			c.Set("chain", global.CONFIG.Contract.DefaultNet) // 默认链
		}
	}
}
