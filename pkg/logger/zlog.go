package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strconv"
	"time"
)

const loggerKey = iota

var (
	RequestId string
	Logger    *zap.Logger
)

// 初始化日志配置

func Init(Service string) {
	filename := ""
	if len(Service) > 0 {
		filename = "logs/" + Service + "/" + time.Now().Format("2006-01-02") + ".log"
	} else {
		filename = "/logs/" + time.Now().Format("2006-01-02") + ".log"
	}
	level := zap.DebugLevel
	NewDevelopmentEncoderConfig := zap.NewDevelopmentEncoderConfig()
	NewDevelopmentEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewTee(
		// 打印在 kafka topic中（伪造的case）
		//zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
		// 打印在控制台
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(NewDevelopmentEncoderConfig),  // json格式日志（ELK渲染收集）
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 打印到控制台和文件
			level,                                                   // 日志级别
		),
		// 打印在文件中
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(NewDevelopmentEncoderConfig),
			zapcore.AddSync(&lumberjack.Logger{
				Filename:   filename, // 日志文件存放目录
				MaxSize:    5,        // 文件大小限制,单位MB
				MaxBackups: 100,      // 最大保留日志文件数量
				MaxAge:     30,       // 日志文件保留天数
				Compress:   false,    // 是否压缩处理
			}),
			level, // 日志级别
		),
	)

	// 开启文件及行号
	development := zap.Development()
	Logger = zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel), // error级别日志，打印堆栈
		development,
	)
	defer func(Logger *zap.Logger) {
		err := Logger.Sync()
		if err != nil {

		}
	}(Logger)
}

// 给指定的context添加字段（关键方法）

func NewContext(ctx *gin.Context, fields ...zapcore.Field) {
	ctx.Set(strconv.Itoa(loggerKey), WithContext(ctx).With(fields...))
}

// 从指定的context返回一个zap实例（关键方法）

func WithContext(ctx *gin.Context) *zap.Logger {
	if ctx == nil {
		return Logger
	}
	l, _ := ctx.Get(strconv.Itoa(loggerKey))
	ctxLogger, ok := l.(*zap.Logger)
	if ok {
		return ctxLogger
	}
	return Logger
}

// Error 当存在错误时记录日志
func Error(err error) {
	if len(RequestId) > 0 {
		Logger.WithOptions(zap.AddCallerSkip(1)).Info("错误日志", zap.String("RequestId", RequestId), zap.Any("详细内容", err))
	} else {
		Logger.WithOptions(zap.AddCallerSkip(1)).Info("错误日志", zap.Any("详细内容", err))
	}
}

// Info 记录日志
func Info(info interface{}) {
	if len(RequestId) > 0 {
		Logger.WithOptions(zap.AddCallerSkip(1)).Info("调试日志", zap.String("RequestId", RequestId), zap.Any("详细内容", info))
	} else {
		Logger.WithOptions(zap.AddCallerSkip(1)).Info("调试日志", zap.Any("详细内容", info))
	}
}
