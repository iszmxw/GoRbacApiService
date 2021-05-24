package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/middlewares/admin"
	"gorbac/routes"
)

// SetupRoute 路由初始化
func SetupRoute(logId string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(admin.DiyCors())
	router.Use(admin.TrackingId(logId))
	// v1 版本
	v1 := router.Group("/v1")
	routes.RegisterWebRoutes(v1)
	return router
}
