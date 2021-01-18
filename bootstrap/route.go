package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/http/middlewares"
	"gorbac/routes"
)

// SetupRoute 路由初始化
func SetupRoute() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.DiyCors())
	// v1 版本
	v1 := router.Group("/v1")
	routes.RegisterWebRoutes(v1)
	return router
}
