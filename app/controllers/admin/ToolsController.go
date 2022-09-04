package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	logger "gorbac/pkg/logger"
	"net/http"
)

type ToolsController struct {
	BaseController
}

// UploadsHandler 上传文件
func (h *ToolsController) UploadsHandler(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	logger.Info(file.Filename)

	dst := fmt.Sprint(`D:\tmp\`+file.Filename)
	// 上传文件到指定的路径
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		logger.Info(err)
		return 
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
