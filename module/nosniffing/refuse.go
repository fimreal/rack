package nosniffing

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 存储被禁止的 IP 地址及其过期时间
type bannedIP struct {
	expiry time.Time
}

var (
	blackUri  = []string{"/wp-admin", "/wp-content", "/.env"}
	bannedIPs = make(map[string]bannedIP) // 存储被禁止的 IP 地址
	mu        sync.Mutex
)

// listBannedIPs 返回当前被禁止的 IP 列表
func listBannedIPs(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	if len(bannedIPs) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No banned IPs found"})
		return
	}

	response := make(map[string]string)
	for ip, ban := range bannedIPs {
		response[ip] = ban.expiry.Format(time.RFC3339) // 格式化到期时间
	}

	c.JSON(http.StatusOK, response)
}

// addBannedIPs 添加封锁 IP
func addBannedIPs(c *gin.Context) {
	ip := c.Param("ip")

	mu.Lock()
	defer mu.Unlock()

	bannedIPs[ip] = bannedIP{expiry: time.Now().Add(10 * time.Minute)} // 将该 IP 加入禁止列表
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("IP %s has been banned", ip)})
}

// unbanIP 解封指定的 IP
func unbanIP(c *gin.Context) {
	ip := c.Param("ip")

	mu.Lock()
	defer mu.Unlock()

	if _, exists := bannedIPs[ip]; exists {
		// 从禁止列表中删除该 IP
		delete(bannedIPs, ip)
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("IP %s has been unbanned", ip)})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("IP %s not found in banned list", ip)})
	}
}

// checkBannedIPs 用于检查请求的 IP 是否被禁止
func checkBannedIPs(c *gin.Context) {
	// 留空解封自己
	requestPath := c.Request.URL.Path
	if requestPath == "/health" || strings.HasPrefix(requestPath, viper.GetString(ID+"_key")) {
		c.Next()
		return
	}

	ip := c.ClientIP() // 获取客户端 IP
	mu.Lock()
	ban, exists := bannedIPs[ip]
	mu.Unlock()

	// 检查是否存在封禁记录
	if exists {
		// 检查是否仍然处于禁止状态
		if time.Now().Before(ban.expiry) {
			c.AbortWithStatus(444)
			return
		}

		// 如果过期，则解除封禁
		mu.Lock()
		delete(bannedIPs, ip)
		mu.Unlock()
	}

	c.Next()
}

// maliciousRequestChecker 检测恶意请求的中间件
func maliciousRequestChecker(c *gin.Context) {
	requestPath := c.Request.URL.Path

	for _, uri := range blackUri {
		if strings.Contains(requestPath, uri) {
			// 调用圈套处理函数禁止访问
			trapHandler(c)
			return
		}
	}

	c.Next()
}

// trapHandler 用来处理恶意请求
func trapHandler(c *gin.Context) {
	ip := c.ClientIP()           // 获取客户端 IP
	duration := 10 * time.Minute // 设置禁止时长，您可以根据需要更改

	mu.Lock()
	bannedIPs[ip] = bannedIP{expiry: time.Now().Add(duration)} // 将该 IP 加入禁止列表
	mu.Unlock()

	// 返回 403 Forbidden 状态码
	c.AbortWithStatus(444)
	ezap.Warnf("Access denied for IP: %s\n", ip)
}

// 清理过期的禁止 IP
func cleanupBannedIPs() {
	for {
		// 每分钟检查一次
		time.Sleep(1 * time.Minute)
		now := time.Now()
		mu.Lock()
		for ip, ban := range bannedIPs {
			// 检查移除过期的禁止 IP
			if now.After(ban.expiry) {
				delete(bannedIPs, ip)
				ezap.Infof("IP %s has been unbanned.", ip)
			}
		}
		mu.Unlock()
	}
}
