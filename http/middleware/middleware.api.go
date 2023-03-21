package middleware

import (
	"github.com/gin-gonic/gin"
)

// 跨域处理
func ApiMonitor() gin.HandlerFunc {
	return func(c *gin.Context) {
		//todo 统计接口
		c.Next()
	}
}
