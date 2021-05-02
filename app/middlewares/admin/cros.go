package admin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorbac/pkg/utils/helpers"
	"gorbac/pkg/utils/logger"
	"net/http"
)

const DefaultHeader = "Tracking-Id"

// Cors 跨域请求 基于gin官方的
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	config.AllowOrigins = []string{"http://localhost:9527", "https://www.gourouting.com"}
	config.AllowCredentials = true
	return cors.New(config)
}

// DiyCors 跨域请求 自定义的
func DiyCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Admin-Token, admin-token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

// TrackingId Gin 具有默认标头的中间件
func TrackingId() gin.HandlerFunc {
	return func(c *gin.Context) {
		tId := c.GetHeader(DefaultHeader)
		// 如果不存在，则生成TrackingID
		if tId == "" {
			tId = helpers.GetUUID()
			c.Header(DefaultHeader, tId)
		}
		// Set in Context
		c.Set(DefaultHeader, tId)
		logger.RequestId = tId
		c.Next()
	}
}
