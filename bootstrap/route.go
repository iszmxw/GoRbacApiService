package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gorbac/app/middlewares/admin"
	"gorbac/routes"
	"net/http"
)

// SetupRoute 路由初始化
func SetupRoute() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(admin.TraceLogger()) // 日志追踪
	router.Use(admin.Cors())        // 跨域
	router.NoRoute(NoResponse)
	router.GET("/", func(context *gin.Context) {
		requestId, _ := context.Get("Tracking-Id")
		context.String(200, "Hello World!\n"+requestId.(string))
	})
	routes.RegisterWebRoutes(router.Group("/"))
	return router
}

func NoResponse(c *gin.Context) {
	//返回404状态码
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404, page not exists!",
	})
}
