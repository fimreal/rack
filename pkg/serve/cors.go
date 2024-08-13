package serve

import (
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var allowedOrigins []string
var mu sync.RWMutex

func Cors() gin.HandlerFunc {
	// 检查是否启用 CORS 限制
	if !viper.GetBool("cors") {
		// 默认允许所有跨域请求
		return cors.Default()
	}

	// 按需配置
	return cors.New(cors.Config{
		AllowAllOrigins: false,
		// AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// 如果配置了 AllowOriginFunc，则 AllowOrigins 不生效
		// AllowOriginFunc: func(origin string) bool {
		// 	// 动态配置
		// 	for _, allowedOrigin := range allowedOrigins {
		// 		if origin == allowedOrigin {
		// 			return true
		// 		}
		// 	}
		// 	return false // 返回 false, 表示不允许其他来源
		// },
		AllowOriginFunc: func(origin string) bool {
			mu.RLock()
			defer mu.RUnlock()
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					return true
				}
			}
			return false
		},
		MaxAge: 12 * 3600,
	})
}

// 动态更新 CORS 允许的来源
// 按需配置前端主机名可能用到
func UpdateAllowedOrigins(c *gin.Context) {
	var newOrigins []string
	if err := c.ShouldBindJSON(&newOrigins); err == nil {
		mu.Lock()
		allowedOrigins = newOrigins
		mu.Unlock()
		c.JSON(200, gin.H{"message": "CORS origins updated"})
	} else {
		c.JSON(400, gin.H{"error": "Invalid input"})
	}
}
