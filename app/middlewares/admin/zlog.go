package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorbac/pkg/config"
	"gorbac/pkg/logger"
	"io/ioutil"
)

const DefaultHeader = "Tracking-Id"

func TraceLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 每个请求生成的请求RequestId具有全局唯一性
		RequestId := ctx.GetHeader(DefaultHeader)
		// 如果不存在，则生成TrackingID
		if RequestId == "" {
			RequestId = uuid.New().String()
			ctx.Header(DefaultHeader, RequestId)
		}
		fmt.Printf("当前请求ID为：%v\n", RequestId)
		ctx.Set(DefaultHeader, RequestId)
		logger.RequestId = RequestId
		logger.NewContext(ctx, zap.String("AppName", config.GetString("app.name")))
		logger.NewContext(ctx, zap.String("AppEnv", config.GetString("app.env")))
		logger.NewContext(ctx, zap.String("RequestId", RequestId))
		// 为日志添加请求的地址以及请求参数等信息
		logger.NewContext(ctx, zap.String("request.ip", ctx.ClientIP()))
		logger.NewContext(ctx, zap.String("request.method", ctx.Request.Method))
		logger.NewContext(ctx, zap.String("request.url", ctx.Request.URL.String()))
		headers, _ := json.Marshal(ctx.Request.Header)
		logger.NewContext(ctx, zap.String("request.headers", string(headers)))
		// 将请求参数json序列化后添加进日志上下文
		data, err := ctx.GetRawData()
		if err != nil {
			logger.Error(err)
		}
		// 很关键,把读过的字节流重新放到body
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		logger.NewContext(ctx, zap.Any("request.params", string(data)))
		logger.Info("=================Start：" + RequestId + "=====================")
		logger.WithContext(ctx).Info("请求信息", zap.Skip())
		logger.Info("=================End：" + RequestId + "=====================")
		ctx.Next()
	}
}
