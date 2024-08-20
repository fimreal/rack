package jwt

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware JWT 认证中间件
func JWTMiddleware(skipPaths []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查当前请求的路径是否在免验证列表中
		for _, path := range skipPaths {
			if strings.HasPrefix(c.Request.URL.Path, path) {
				c.Next() // 跳过 Token 验证
				return
			}
		}

		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "请求缺少 Token"})
			c.Abort()
			return
		}

		// 解析和验证 token
		valid, claims, err := DecodeToken(tokenString, "access"+Secret)
		if err != nil || !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 Token"})
			c.Abort()
			return
		}

		// 将用户 ID 或其他信息存储到上下文中，供后续使用
		id := (*claims)["id"].(uint)
		c.Set("id", int64(id)) // 设置一个新的上下文键值对

		c.Next() // 继续处理请求
	}
}
