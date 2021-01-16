package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gorbac/routes"
)

// SetupRoute 路由初始化
func SetupRoute() *gin.Engine {
	router := gin.Default()
	routes.RegisterWebRoutes(router)

	return router
}
