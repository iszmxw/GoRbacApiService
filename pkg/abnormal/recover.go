package abnormal

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorbac/pkg/logger"
	"net"
	"os"
	"runtime/debug"
	"strings"
)

func Stack(description string, c *gin.Context) {
	if r := recover(); r != nil {
		var brokenPipe = false              // <--网络连接是否断开的判断条件
		if ne, ok := r.(*net.OpError); ok { // <--这里对brokenPipe进行设置
			if se, oks := ne.Err.(*os.SyscallError); oks {
				if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
					brokenPipe = true
				}
			}
		}

		if c != nil && brokenPipe { // <--如果连接已经中断，我们只需要简单的记录并中断处理
			logger.Error(errors.New("网络连接中断"))
			//NoticeMail.Notice("msgNetworkAnomaly", map[string]interface{}{
			//	"title":       "网络连接中断",
			//	"c":           c,
			//	"r":           r,
			//	"description": description,
			//	"errInfo":     nil,
			//})
			c.Abort()
		}
		// 收集错误堆栈信息 异常
		errInfo := string(debug.Stack())
		logger.Error(errors.New(fmt.Sprintf(`描述: %v`, description)))
		logger.Error(errors.New(fmt.Sprintf(`Recover: %v`, r)))
		logger.Error(errors.New(fmt.Sprintf(`详细堆栈错误信息: %v`, errInfo)))
		// 网络连接中断就不发送了
		if brokenPipe == false {
			//NoticeMail.Notice("msgNetworkAnomaly", map[string]interface{}{
			//	"title":       "接口内部异常",
			//	"c":           c,
			//	"r":           r,
			//	"description": description,
			//	"errInfo":     errInfo,
			//})
		}
	}
}
