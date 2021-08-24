package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/middlewares/admin"
	"gorbac/routes"
)

// SetupRoute 路由初始化
func SetupRoute() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(admin.DiyCors())
	router.Use(admin.TrackingId())
	// v1 版本
	v1 := router.Group("/v1")
	router.GET("/", func(context *gin.Context) {
		requestId, _ := context.Get("Tracking-Id")
		context.String(200, "Hello World!\n"+requestId.(string))
	})
	routes.RegisterWebRoutes(v1)
	return router
}
