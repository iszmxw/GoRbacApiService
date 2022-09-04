package pprof

import (
	"fmt"
	"gorbac/pkg/config"
	"gorbac/pkg/logger"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"runtime/debug"
)

// 开启 pprof 分析

func Debug(addr string) {
	// pprof 分析
	if config.GetBool("app.pprof") {
		runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
		runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪
		go func() {
			defer func() {
				if r := recover(); r != nil {
					title := fmt.Sprintf(`【%v】捕捉到异常：当前端口【%v】`, "pprof", addr)
					cutContent := r
					errInfo := string(debug.Stack())
					logger.Info("pprof=" + addr)
					logger.Info(cutContent)
					logger.Info(errInfo)
					logger.Info(map[string]interface{}{
						"title":   title,
						"content": errInfo,
					})
				}
			}()
			// 启动一个 http server，注意 pprof 相关的 handler 已经自动注册过了
			if err := http.ListenAndServe(addr, nil); err != nil {
				logger.Error(err)
			}
		}()
	}
}
