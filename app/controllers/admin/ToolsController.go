package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorbac/pkg/utils/logger"
	"net/http"
)

type ToolsController struct {
}

// UploadsHandler 上传文件
func (h *ToolsController) UploadsHandler(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	logger.LogInfo(file.Filename)

	dst := fmt.Sprint(`D:\tmp\`+file.Filename)
	// 上传文件到指定的路径
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		logger.LogInfo(err)
		return 
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
