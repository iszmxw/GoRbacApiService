package echo

import (
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorbac/pkg/logger"
	"gorbac/pkg/utils/helpers"
)

// Rjson 成功返回封装 参数 data interface{} 类型为可接受任意类型
func Rjson(c *gin.Context, result interface{}, msg string, code int) {
	reqId, _ := c.Get("Tracking-Id")
	rdata := cmap.New().Items()
	rdata["reqId"] = reqId
	rdata["code"] = code
	if result != nil {
		rdata["data"] = result
	}
	if len(msg) > 0 {
		rdata["msg"] = msg
	} else {
		rdata["msg"] = "success."
	}
	c.JSON(200, rdata)
	return
}

// Error  错误返回封装
func Error(c *gin.Context, code string, lastMsg string) {
	var logInfo []zapcore.Field
	if len(logger.RequestId) > 0 {
		logInfo = append(logInfo, zap.Any("RequestId", logger.RequestId))
	}
	code, firstMsg := GetCode(code) // 最终code
	fullMsg := firstMsg + lastMsg   // 最终msg
	logInfo = append(logInfo, zap.Any(code, fullMsg))
	logger.Logger.WithOptions(zap.AddCallerSkip(1)).Info("返回错误", logInfo...)
	Rjson(c, nil, fullMsg, helpers.StringToInt(code))
}

// Success  错误返回封装
func Success(c *gin.Context, result interface{}, msg string) {
	logger.Logger.WithOptions(zap.AddCallerSkip(1)).Info("成功返回", zap.Any("返回数据", result))
	Rjson(c, result, msg, 1)
}
