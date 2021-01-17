package bootstrap

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorbac/routes"
)

// SetupRoute 路由初始化
func SetupRoute() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	// v1 版本
	v1 := router.Group("/v1")
	routes.RegisterWebRoutes(v1)
	return router
}
